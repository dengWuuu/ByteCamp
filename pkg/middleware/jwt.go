package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"douyin/cmd/api/handlers"
	"douyin/cmd/user/pack"
	"douyin/dal/db"
	"douyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func GetUserIdFromJwtToken(ctx context.Context, c *app.RequestContext) uint {
	claims, err := JwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		unauthorized(ctx, c, http.StatusUnauthorized, JwtMiddleware.HTTPStatusMessageFunc(err, ctx, c))
		return 0
	}
	userMap := claims[jwt.IdentityKey].(map[string]interface{})
	userId := uint(userMap["ID"].(float64))
	return userId
}

func unauthorized(ctx context.Context, c *app.RequestContext, code int, message string) {
	c.Header("WWW-Authenticate", "JWT realm="+JwtMiddleware.Realm)
	if !JwtMiddleware.DisabledAbort {
		c.Abort()
	}

	JwtMiddleware.Unauthorized(ctx, c, code, message)
}

func JwtMiddlewareFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 验证服务端token有么有过期
		userId := GetUserIdFromJwtToken(ctx, c)
		isNil := db.UserRedis.Get(ctx, "user:"+strconv.Itoa(int(userId)))
		if isNil.Val() == "" {
			unauthorized(ctx, c, http.StatusUnauthorized, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrExpiredToken, ctx, c))
		}
		// 验证客户端token有没有过期
		claims, err := JwtMiddleware.GetClaimsFromJWT(ctx, c)
		if err != nil {
			unauthorized(ctx, c, http.StatusUnauthorized, JwtMiddleware.HTTPStatusMessageFunc(err, ctx, c))
			return
		}

		switch v := claims["exp"].(type) {
		case nil:
			unauthorized(ctx, c, http.StatusBadRequest, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrMissingExpField, ctx, c))
			return
		case float64:
			if int64(v) < JwtMiddleware.TimeFunc().Unix() {
				unauthorized(ctx, c, http.StatusUnauthorized, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrExpiredToken, ctx, c))
				return
			}
		case json.Number:
			n, err := v.Int64()
			if err != nil {
				unauthorized(ctx, c, http.StatusBadRequest, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrWrongFormatOfExp, ctx, c))
				return
			}
			if n < JwtMiddleware.TimeFunc().Unix() {
				unauthorized(ctx, c, http.StatusUnauthorized, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrExpiredToken, ctx, c))
				return
			}
		default:
			JwtMiddleware.Unauthorized(ctx, c, http.StatusBadRequest, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrWrongFormatOfExp, ctx, c))
		}

		c.Set("JWT_PAYLOAD", claims)
		identity := JwtMiddleware.IdentityHandler(ctx, c)

		if identity != nil {
			c.Set(JwtMiddleware.IdentityKey, identity)
		}

		if !JwtMiddleware.Authorizator(identity, ctx, c) {
			unauthorized(ctx, c, http.StatusForbidden, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrForbidden, ctx, c))
			return
		}

		c.Next(ctx)
	}
}

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:            "DouYin JwtUtils",
		SigningAlgorithm: "HS256",
		Key:              []byte("secret key"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		TokenLookup:      "header: Authorization, query: token, cookie: jwt, param: token, form: token",
		TokenHeadName:    "Bearer",

		// 构造登录成功的返回请求
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			// 从token中获取用户id
			Token, err := JwtMiddleware.ParseTokenString(token)
			claims := jwt.ExtractClaimsFromToken(Token)
			userMap, _ := claims[IdentityKey].(map[string]interface{})
			// 将token同时存进redis
			userId := uint(userMap["ID"].(float64))

			db.UserRedis.Set(ctx, "user"+":"+strconv.Itoa(int(userId)), token, time.Hour*24)
			if err != nil {
				hlog.Fatalf("不能从Jwt中获取claims")
				c.JSON(10086, "登录请求响应失败")
			}
			c.JSON(http.StatusOK, utils.H{
				"status_code": 0,
				"token":       token,
				"user_id":     userId,
				"status_msg":  "success",
				"expire_time": expire.Format(time.RFC3339),
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				UserName string `json:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
				Password string `json:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			}
			loginStruct.UserName = c.Query("username")
			loginStruct.Password = c.Query("password")
			if len(loginStruct.UserName) == 0 || len(loginStruct.Password) == 0 {
				handlers.SendResponse(c, pack.BuildUserRegisterResp(errno.ErrBind))
				return nil, nil
			}

			users, err := db.CheckUser(loginStruct.UserName, loginStruct.Password)
			if err != nil {
				return nil, err
			}
			if len(users) == 0 {
				return nil, errors.New("user already exists or wrong password")
			}

			return users[0], nil
		},
		IdentityKey: IdentityKey,
		// 登录成功后获取请求中jwt中存储用户的id
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[IdentityKey].(map[string]interface{})
		},
		// 登录成功 token中放入userId信息
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*db.User); ok {
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "JwtUtils biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}

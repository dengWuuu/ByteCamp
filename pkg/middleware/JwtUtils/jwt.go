package JwtUtils

import (
	"context"
	"douyin/dal/db"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func GetUserIdFromJwtToken(ctx context.Context, c *app.RequestContext) uint {
	claims := jwt.ExtractClaims(ctx, c)
	userMap := claims[jwt.IdentityKey].(map[string]interface{})
	userId := uint(userMap["ID"].(float64))
	return userId
}

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:            "DouYin JwtUtils",
		SigningAlgorithm: "HS256",
		Key:              []byte("secret key"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		TokenLookup:      "header: Authorization, query: token, cookie: JwtUtils",
		TokenHeadName:    "Bearer",
		//构造登录成功的返回请求
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			//从token中获取用户id
			Token, err := JwtMiddleware.ParseTokenString(token)
			claims := jwt.ExtractClaimsFromToken(Token)
			userMap, _ := claims[IdentityKey].(jwt.MapClaims)
			if err != nil {
				hlog.Fatalf("不能从Jwt中获取claims")
				c.JSON(10086, "登录请求响应失败")
			}
			c.JSON(http.StatusOK, utils.H{
				"status_code": 0,
				"token":       token,
				"user_id":     userMap["ID"],
				"status_msg":  "success",
				"expire_time": expire.Format(time.RFC3339),
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				UserName string `json:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
				Password string `json:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				hlog.Infof("jwt认证用户数据绑定失败")
				return nil, err
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
		//登录成功后获取请求中jwt中存储用户的id
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[IdentityKey].(map[string]interface{})
		},
		//登录成功 token中放入userId信息
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

package redis

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"douyin/dal/db"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
)

const prefix = "user:Id:"

func GetUsersFromRedis(ctx context.Context, userIds []uint) []*db.User {
	userList := make([]*db.User, len(userIds))
	// 使用管道命令
	pipelined := db.UserRedis.Pipeline()
	for i := 0; i < len(userIds); i++ {
		pipelined.Get(ctx, prefix+strconv.Itoa(int(userIds[i])))
	}
	res, err := pipelined.Exec(ctx)
	if err != nil {
		klog.Fatal("管道命令失败")
	}
	for index, cmdRes := range res {
		cmd, ok := cmdRes.(*redis.StringCmd)
		if !ok {
			klog.Fatal("redis相关强转失败")
		}
		bytes, err := cmd.Bytes()
		if err != nil {
			klog.Fatal("redis获取用户失败,获取字节码数组失败")
			userList[index] = nil
			continue
		}
		user := new(db.User)
		err = json.Unmarshal(bytes, user)
		if err != nil {
			klog.Fatalf("redis中获取用户信息后反序列化失败")
			return nil
		}
		userList[index] = user
	}
	return userList
}

func PutUserToRedis(ctx context.Context, user *db.User) {
	marshal, err := json.Marshal(user)
	if err != nil {
		klog.Fatalf("将user放入redis时序列化失败")
	}

	result := db.UserRedis.Set(ctx, prefix+strconv.Itoa(int(user.ID)), marshal, time.Hour*24*180)
	s, err := result.Result()
	klog.Info("redis放入user信息" + s)
}

func LoadUserFromMysqlToRedis(ctx context.Context) {
	user := db.GetAllUser()
	pipelined := db.UserRedis.Pipeline()
	for i := 0; i < len(user); i++ {
		marshal, err := json.Marshal(user[i])
		if err != nil {
			klog.Fatalf("将user放入redis时序列化失败")
		}
		pipelined.Set(ctx, prefix+strconv.Itoa(int(user[i].ID)), marshal, time.Hour*24*180)
	}
	_, err := pipelined.Exec(ctx)
	if err != nil {
		klog.Fatalf("LoadUserFromMysqlToRedis管道命令执行失败")
	}
}

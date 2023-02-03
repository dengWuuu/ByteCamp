package test

import (
	"context"
	"douyin/dal/db"
	"testing"
	"time"
)

func TestUserRedis(t *testing.T) {
	ctx := context.Background()
	db.Init("../config")
	db.UserRedis.Set(ctx, "user:1", "akjdkasdhfkljlsadhf", time.Minute*10)
}

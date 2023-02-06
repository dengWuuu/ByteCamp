package test

import (
	"context"
	"douyin/cmd/relation/service"
	"douyin/dal/db"
	"testing"
)

func TestIsFollowing(t *testing.T) {
	db.Init("../config")
	ctx := context.Background()
	res, _ := service.IsFollowing(ctx, 1, 2)
	t.Log(res)
	res, _ = service.IsFollowing(ctx, 1, 5)
	t.Log(res)
	res, _ = service.IsFollowing(ctx, 1, 10)
	t.Log(res)
	res, _ = service.IsFollowing(ctx, 23, 24)
	t.Log(res)
}

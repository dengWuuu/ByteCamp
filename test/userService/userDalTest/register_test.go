package userService

import (
	"douyin/dal/db"
	"fmt"
	"testing"
)

func TestRpcRegistry(t *testing.T) {
	name, err := db.GetUsersByUserName("wdw")
	if err != nil {
		return
	}
	fmt.Println(name)
}

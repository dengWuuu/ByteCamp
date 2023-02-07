package userService

import (
	"fmt"
	"testing"

	"douyin/dal/db"
)

func TestRpcRegistry(t *testing.T) {
	name, err := db.GetUsersByUserName("wdw")
	if err != nil {
		return
	}
	fmt.Println(name)
}

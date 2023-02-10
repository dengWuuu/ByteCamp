package test

import (
	"douyin/dal/db"
	"testing"
)

func TestMigrate(t *testing.T) {
	db.Init("../config")
	db.DB.Migrator().DropColumn(&db.Comment{}, "creat_time")
}

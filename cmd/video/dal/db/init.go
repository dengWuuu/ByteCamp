package db

import (
	"log"
	"os"
	"time"

	"douyin/cmd/video/config"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB
var dbErr error

func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	DB, dbErr = gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{Logger: newLogger})
	if dbErr != nil {
		log.Fatal("mysql连接失败")
	}
	err := DB.Use(
		dbresolver.Register(dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200))
	if err != nil {
		hlog.Fatalf("数据库连接池失败")
		return
	}
}

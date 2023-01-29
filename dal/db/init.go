package db

import (
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var Redis *redis.Client
var dbErr error

func Init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置文件失败")
		return
	}
	InitMysql()
	InitRedis()
}

func InitMysql() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	dsn := viper.GetString("mysql.dsn")
	DB, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if dbErr != nil {
		log.Fatal("mysql连接失败")
	}
}

func InitRedis() {
	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	poolSize := viper.GetInt("redis.poolSize")
	minConns := viper.GetInt("redis.minConns")

	Redis = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		PoolSize:     poolSize,
		MinIdleConns: minConns,
	})
}

package db

import (
	"log"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var ExpireTime time.Duration

var DB *gorm.DB
var dbErr error

// FollowingRedis relation部分redis客户端
var FollowingRedis *redis.Client
var FollowersRedis *redis.Client
var FriendsRedis *redis.Client
var UserRedis *redis.Client
var CommentRedis *redis.Client

func Init(configPath string) {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
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

func InitRedis() {
	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	poolSize := viper.GetInt("redis.poolSize")
	minConns := viper.GetInt("redis.minConns")

	ExpireTime = viper.GetDuration("redis.exipretime")

	hlog.Info("followingdb:%v,followerdb:%v,friendsdb:%v\n", viper.GetInt("redis.followingdb"), viper.GetInt("redis.followersdb"), viper.GetInt("redis.friendsdb"))

	FollowingRedis = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		PoolSize:     poolSize,
		MinIdleConns: minConns,
		DB:           viper.GetInt("redis.followingdb"),
	})

	FollowersRedis = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		PoolSize:     poolSize,
		MinIdleConns: minConns,
		DB:           viper.GetInt("redis.followersdb"),
	})

	FriendsRedis = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		PoolSize:     poolSize,
		MinIdleConns: minConns,
		DB:           viper.GetInt("redis.friendsdb"),
	})
	UserRedis = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		PoolSize:     poolSize,
		MinIdleConns: minConns,
		DB:           viper.GetInt("redis.userdb"),
	})
	CommentRedis = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		PoolSize:     poolSize,
		MinIdleConns: minConns,
		DB:           viper.GetInt("redis.commentdb"),
	})
}

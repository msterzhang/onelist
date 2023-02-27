package database

import (
	"log"
	"os"
	"time"

	"github.com/msterzhang/onelist/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func NewDb() *gorm.DB {
	return db
}

func InitDb() error {
	var err error
	dia := sqlite.Open(config.DbName + ".db")
	if config.DBDRIVER == "mysql" {
		dia = mysql.Open(config.DBURL)
	}
	db, err = gorm.Open(dia, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold:             time.Second,   // 慢 SQL 阈值
				LogLevel:                  logger.Silent, // 日志级别
				IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,         // 禁用彩色打印
			},
		),
	})
	if err != nil {
		log.Fatal("数据库打开失败!")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("连接数据库失败!")
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

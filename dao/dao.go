package dao

import (
	"ginProject/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open(mysql.Open(config.MysqlDb), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接失败: connect error", err)
		panic(err)
	}

	// 设置连接池
	sqlDB, err := Db.DB()
	if err != nil {
		log.Println("获取底层数据库对象失败:", err)
		panic(err)
	} else {
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
}

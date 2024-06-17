package dao

import (
	"ginProject/config"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Gormv2 *gorm.DB
	err    error
	once   sync.Once
)

// 单例模式
func GetDb() *gorm.DB {
	once.Do(func() {
		Gormv2, err = gorm.Open(mysql.Open(config.MysqlDb), &gorm.Config{})
		if err != nil {
			log.Println("数据库连接失败: connect error", err)
			panic(err)
		}

		// 设置连接池
		sqlDB, err := Gormv2.DB()
		if err != nil {
			log.Println("获取底层数据库对象失败:", err)
			panic(err)
		} else {
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(time.Hour)
		}
	})
	return Gormv2
}

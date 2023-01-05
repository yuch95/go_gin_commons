package db_dao

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func CreateGormDb(username, pwd, host string, port int, dbname string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, username, pwd, dbname, port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func SetDbConn(gorm *gorm.DB, maxIdleConn, maxOpenConn int, maxIdleTime, maxLifetime time.Duration) {
	db, err := gorm.DB()
	if err != nil {
		panic(err)
	}
	// 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(maxIdleConn)
	// 设置打开数据库连接的最大数量
	db.SetMaxOpenConns(maxOpenConn)
	// 设置连接可能空闲的最长时间
	db.SetConnMaxIdleTime(maxIdleTime)
	// 设置连接可重复使用的最长时间
	db.SetConnMaxLifetime(maxLifetime)
}

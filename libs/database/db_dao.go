package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var GormDb *gorm.DB

func NewGormDb(username, pwd, host string, port int, dbname string, insertSize int) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, username, pwd, dbname, port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{CreateBatchSize: insertSize})
}

func SetDbConn(gormDb *gorm.DB, maxIdleConn, maxOpenConn int, maxIdleTime, maxLifetime time.Duration) {
	db, err := gormDb.DB()
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

type DbDao struct {
	Db *gorm.DB
}

func (d *DbDao) Insert(obj any) error {
	result := d.Db.Create(obj)
	return result.Error
}

func (d *DbDao) InsertMany(obj []any) error {
	result := d.Db.Create(obj)
	return result.Error
}

func (d *DbDao) Update(obj any, queryParam *QueryParams, updateData map[string]any) error {
	query := d.Db.Where(queryParam.QueryData)
	if queryParam.Entities != nil {
		query = query.Select(queryParam.Entities)
	}
	if queryParam.OrderFields != "" {
		query = query.Order(queryParam.OrderFields)
	}
	if queryParam.DistinctFields != nil {
		query.Distinct(queryParam.DistinctFields)
	}
	//TODO implement me
	panic("implement me")
}

func (d *DbDao) UpdateInsert(obj any, queryParam *QueryParams, updateData map[string]any) error {
	//TODO implement me
	panic("implement me")
}

func (d *DbDao) QueryOne(obj any, queryParam *QueryParams) {
	//TODO implement me
	panic("implement me")
}

func (d *DbDao) QueryAll(obj any, queryParam *QueryParams) {
	//TODO implement me
	panic("implement me")
}

func (d *DbDao) PageQuery(obj any, queryParam *QueryParams) {
	//TODO implement me
	panic("implement me")
}

func (d *DbDao) Delete(obj any, queryParam *QueryParams) {
	//TODO implement me
	panic("implement me")
}

func (d *DbDao) DeletePhysical(obj any, queryParam *QueryParams) {
	//TODO implement me
	panic("implement me")
}

func (d *DbDao) Commit() {
	//TODO implement me
	panic("implement me")
}

func (d *DbDao) Rollback() {
	//TODO implement me
	panic("implement me")
}

func init() {
	GormDb, _ = NewGormDb("postgres", "postgres", "localhost", 9920, "postgres", 1000)
}

package mysql

import (
	"context"
	"fmt"
	"github.com/adnpa/IM/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func init() {
	var err error

	cfg := config.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.DBUserName, cfg.DBPassword, cfg.DBAddress, cfg.DBDatabaseName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AllowGlobalUpdate = true
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(config.Config.Mysql.DBMaxOpenConns)
	sqlDB.SetMaxIdleConns(config.Config.Mysql.DBMaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Config.Mysql.DBMaxLifeTime) * time.Second)

	//migration()
	return
}

// Close 关闭数据库连接
func Close() {
	sqldb, _ := db.DB()
	sqldb.Close()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
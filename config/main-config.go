package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// dbUserMysql = "cicuser"
	// dbPassMysql = "Ax12345678"
	dbUserMysql = "admin_root"
	dbPassMysql = "ofXiCHCg8D"

	db = [2]string{"localhost", "admin_icesystem"}
	//db = [2]string{"192.168.60.189", "cicsupport"}
	//db = [2]string{"192.168.1.104", "cicsupport"}

	dsn1 = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUserMysql, dbPassMysql,
		db[0], db[1])
)

// type SqlLogger struct {
// 	logger.Interface
// }

func SetupDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn1), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		// DryRun: true,
		//Logger:                 &SqlLogger{},
	})
	if err != nil {
		panic("Failed to Connect Database")
	}
	//  db.Migrator().CreateTable(User{})
	return db
}

func CloseDBConn(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close Database.")
	}
	dbSQL.Close()
}

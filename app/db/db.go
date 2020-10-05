package db

import (
	"sync"

	"github.com/evandroferreiras/gopher-city-bank/app/common/envvar"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

// newDB creates newDB instance of GORM
func newDB() *gorm.DB {
	dsn := envvar.MySQLDSN()
	logrus.Debug(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panic("error when trying to connect Database")
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Panic("error when trying to connect SQL Database")
		return nil
	}
	sqlDB.SetMaxIdleConns(3)

	return db
}

// DB is a singleton for ORM instance
func DB() *gorm.DB {
	once.Do(func() {
		db = newDB()
	})
	return db
}

// AutoMigrate creates all needed tables
func AutoMigrate() {
	db := DB()
	err := db.AutoMigrate(
		&model.Account{},
		&model.Transfer{},
	)
	if err != nil {
		logrus.Panic("error when trying to AutoMigrate tables/collections")
	}
}

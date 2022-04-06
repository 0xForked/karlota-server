package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

const (
	mysqlDriver  = "mysql"
	sqliteDriver = "sqlite"
)

func (cfg Config) InitDbConn() {
	log.Println("Trying to open database connection . . . .")
	conn, err := openConnection(cfg)

	if err != nil {
		log.Panicln(fmt.Sprintf(
			"DATABASE_ERROR: %s",
			err.Error()))
	}

	log.Println(fmt.Sprintf("Database connected with %s driver . . . .", cfg.GetDbDriver()))
	setConnection(conn)
}

func openConnection(cfg Config) (db *gorm.DB, err error) {
	var driver gorm.Dialector

	switch cfg.GetDbDriver() {
	case sqliteDriver:
		driver = sqlite.Open(cfg.GetDbDsnUrl())
		break
	case mysqlDriver:
		driver = mysql.Open(cfg.GetDbDsnUrl())
		break
	default:
		driver = mysql.Open(cfg.GetDbDsnUrl())
		break
	}

	return gorm.Open(driver, &gorm.Config{})
}

func setConnection(conn *gorm.DB) {
	db = conn
}

func (cfg Config) GetDbConn() *gorm.DB {
	return db
}

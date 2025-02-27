package config

import (
	"github.com/aasumitro/karlota/internal/api/domain"
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
		log.Panicf("DATABASE_ERROR: %s", err.Error())
	}

	log.Printf("Database connected with %s driver . . . .", cfg.GetDbDriver())
	setConnection(conn)

	log.Println("Auto migrate tables . . . .")
	runMigration(conn)
}

func openConnection(cfg Config) (db *gorm.DB, err error) {
	var driver gorm.Dialector

	switch cfg.GetDbDriver() {
	case sqliteDriver:
		driver = sqlite.Open(cfg.GetDbDsnUrl())
	case mysqlDriver:
		driver = mysql.Open(cfg.GetDbDsnUrl())
	default:
		log.Panicf("DATABASE_ERROR: Database driver not supported!")
	}

	return gorm.Open(driver, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
}

func setConnection(conn *gorm.DB) {
	db = conn
}

func runMigration(conn *gorm.DB) {
	domain.User{}.Migrate(conn)
}

func (cfg Config) GetDbConn() *gorm.DB {
	return db
}

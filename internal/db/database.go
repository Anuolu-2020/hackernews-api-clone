package db

import (
	"log"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Anuolu-2020/hackernews-api-clone/pkg/env"
)

var Db *gorm.DB

func InitDB() {
	dbUrl := pkg.GetEnv("DB_URL")

	conn, _ := url.Parse(dbUrl)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"

	db, err := gorm.Open(postgres.Open(conn.String()), &gorm.Config{})
	if err != nil {
		log.Panicf("Failed to connect to db: %v", err)
	}

	Db = db
}

func Close() error {
	sqlDb, err := Db.DB()
	if err != nil {
		log.Fatalf("Failed to convert sql.DB interface: %v", err)
	}
	return sqlDb.Close()
}

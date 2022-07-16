package database

import (
	"database/sql"
	"fmt"
	"reflect"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
	SSLMode  string
}

func NewPostgres(cfg Config) *sql.DB {
	db := mustOpen("postgres",
		fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DB, cfg.SSLMode))

	mustPing(db)
	return db
}

func GetDriverName(db *sql.DB) string {
	driver := db.Driver()

	for _, driverName := range sql.Drivers() {
		db, _ := sql.Open(driverName, "")
		if db == nil {
			continue
		}

		if reflect.TypeOf(driver) == reflect.TypeOf(db.Driver()) {
			return driverName
		}
	}

	return ""
}

func mustOpen(driverName, connectionString string) *sql.DB {
	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		panic(err)
	}
	return db
}

func mustPing(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

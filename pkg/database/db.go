package database

import (
	"database/sql"
	"reflect"
)

type ConfigDB struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
	SSLMode  string
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

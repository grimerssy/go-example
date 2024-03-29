package database

import (
	"database/sql"
	"fmt"
)

func NewPostgres(cfg ConfigDB) (*sql.DB, func()) {
	db := mustOpen("postgres",
		fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DB, cfg.SSLMode))

	mustPing(db)
	cleanup := func() {
		db.Close()
	}
	return db, cleanup
}

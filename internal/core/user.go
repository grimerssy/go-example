package core

import (
	"database/sql/driver"
)

const UserTable = "users"

var (
	UserRowNames = []string{"id", "name", "password", "count"}
	UserRowMocks = []driver.Value{1, "n", "p", 1}
)

type User struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Count    int    `db:"count"`
}

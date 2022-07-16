package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/consts"
	"github.com/grimerssy/go-example/pkg/database"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	if reflect.ValueOf(db).IsNil() {
		panic("db cannot be nil")
	}
	return &UserRepository{
		db: sqlx.NewDb(db, database.GetDriverName(db)),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *core.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	query := fmt.Sprintf(`
INSERT INTO %s (name, password)
VALUES (:name, :password);
`, core.UserTable)
	_, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id int64) (*core.User, error) {
	user := new(core.User)
	query := fmt.Sprintf(`
SELECT %s FROM %s
WHERE id = $1
LIMIT 1;
`, strings.Join(core.UserRowNames, ", "), core.UserTable)
	err := r.db.GetContext(ctx, user, query, id)
	if err != nil {
		return nil, consts.ErrUserNotFound
	}
	return user, nil
}

func (r *UserRepository) GetUserByName(ctx context.Context, name string) (*core.User, error) {
	user := new(core.User)
	query := fmt.Sprintf(`
SELECT %s FROM %s
WHERE name = $1
LIMIT 1;
`, strings.Join(core.UserRowNames, ", "), core.UserTable)
	err := r.db.GetContext(ctx, user, query, name)
	if err != nil {
		return nil, consts.ErrUserNotFound
	}
	return user, nil
}

func (r *UserRepository) UpdateUserCount(ctx context.Context, user *core.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	query := fmt.Sprintf(`
UPDATE %s
SET count = :count
WHERE id = :id
`, core.UserTable)
	_, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/grimerssy/go-example/internal/core"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewUserRepository", func() {
	var (
		ur *UserRepository

		db *sql.DB

		cleanup func()

		test func()
	)

	JustBeforeEach(func() {
		test = func() {
			ur = NewUserRepository(db)
		}
	})

	When("db is nil", func() {
		BeforeEach(func() {
			db = nil
			cleanup = func() {}
		})

		It("panics", func() {
			Expect(test).To(Panic())
		})
		It("returns nil", func() {
			Expect(ur).To(BeNil())
		})
	})

	When("db is not nil", func() {
		BeforeEach(func() {
			var err error
			db, _, err = sqlmock.New()
			Expect(db).NotTo(BeNil())
			Expect(err).NotTo(HaveOccurred())

			cleanup = func() {
				db.Close()
			}
		})

		It("does not panic", func() {
			Expect(test).NotTo(Panic())
		})
		It("returns non-nil *userRepository", func() {
			Expect(ur).NotTo(BeNil())
		})
	})

	AfterEach(func() {
		cleanup()
	})
})

var _ = Describe("UserRepository", func() {
	var (
		ur *UserRepository

		dbMock sqlmock.Sqlmock
	)

	BeforeEach(func() {
		ur, dbMock = mockUserRepository()
	})

	Describe("CreateUser", func() {
		var (
			ctx  context.Context
			user *core.User

			err error

			insertIntoOK = func() {
				dbMock.ExpectExec(
					fmt.Sprintf(`INSERT INTO %s (.+) VALUES (.+);`,
						core.UserTable)).
					WithArgs(user.Name, user.Password).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}
		)

		BeforeEach(func() {
			ctx = context.Background()
			user = &core.User{}
		})

		JustBeforeEach(func() {
			err = ur.CreateUser(ctx, user)
		})

		When("user is nil", func() {
			BeforeEach(func() {
				user = nil
			})

			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("INSERT INTO query fails", func() {
			BeforeEach(func() {
				dbMock.ExpectExec(
					fmt.Sprintf(`INSERT INTO %s (.+) VALUES (.+);`,
						core.UserTable)).
					WithArgs(user.Name, user.Password).
					WillReturnError(errors.New(""))
			})

			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("no errors are met", func() {
			BeforeEach(func() {
				insertIntoOK()
			})

			It("succeeds", func() {
				Expect(err).To(Succeed())
			})
		})

		AfterEach(func() {
			Expect(dbMock.ExpectationsWereMet()).NotTo(HaveOccurred())
		})
	})

	Describe("GetUserById", func() {
		var (
			ctx context.Context
			id  int64

			user *core.User
			err  error

			selectOK = func() {
				dbMock.ExpectQuery(
					fmt.Sprintf("SELECT .+ FROM %s WHERE id = .+ LIMIT 1;",
						core.UserTable)).
					WithArgs(id).
					WillReturnRows(sqlmock.NewRows(core.UserRowNames).
						AddRow(core.UserRowMocks...))
			}
		)

		BeforeEach(func() {
			ctx = context.Background()
			id = 0
		})

		JustBeforeEach(func() {
			user, err = ur.GetUserById(ctx, id)
		})

		When("SELECT query fails", func() {
			BeforeEach(func() {
				dbMock.ExpectQuery(
					fmt.Sprintf(`SELECT .+ FROM %s WHERE id = .+ LIMIT 1;`,
						core.UserTable)).
					WithArgs(id).
					WillReturnError(errors.New(""))
			})

			It("returns nil *User", func() {
				Expect(user).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns ErrUserNotFound", func() {
				Expect(errors.Is(err, core.ErrUserNotFound)).To(BeTrue())
			})
		})

		When("user is not found", func() {
			BeforeEach(func() {
				dbMock.ExpectQuery(
					fmt.Sprintf(`SELECT .+ FROM %s WHERE id = .+ LIMIT 1;`,
						core.UserTable)).
					WithArgs(id).
					WillReturnRows(sqlmock.NewRows(core.UserRowNames))
			})

			It("returns nil *User", func() {
				Expect(user).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns ErrUserNotFound", func() {
				Expect(errors.Is(err, core.ErrUserNotFound)).To(BeTrue())
			})
		})

		When("no errors are met", func() {
			BeforeEach(func() {
				selectOK()
			})

			It("returns non-nil *User", func() {
				Expect(user).NotTo(BeNil())
			})
			It("succeeds", func() {
				Expect(err).To(Succeed())
			})
		})
	})

	Describe("GetUserByName", func() {
		var (
			ctx  context.Context
			name string

			user *core.User
			err  error

			selectOK = func() {
				dbMock.ExpectQuery(
					fmt.Sprintf("SELECT .+ FROM %s WHERE name = .+ LIMIT 1;",
						core.UserTable)).
					WithArgs(name).
					WillReturnRows(sqlmock.NewRows(core.UserRowNames).
						AddRow(core.UserRowMocks...))
			}
		)

		BeforeEach(func() {
			ctx = context.Background()
			name = ""
		})

		JustBeforeEach(func() {
			user, err = ur.GetUserByName(ctx, name)
		})

		When("SELECT query fails", func() {
			BeforeEach(func() {
				dbMock.ExpectQuery(
					fmt.Sprintf(`SELECT .+ FROM %s WHERE name = .+ LIMIT 1;`,
						core.UserTable)).
					WithArgs(name).
					WillReturnError(errors.New(""))
			})

			It("returns nil *User", func() {
				Expect(user).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns ErrUserNotFound", func() {
				Expect(err).NotTo(Succeed())
				Expect(errors.Is(err, core.ErrUserNotFound)).To(BeTrue())
			})
		})

		When("user is not found", func() {
			BeforeEach(func() {
				dbMock.ExpectQuery(
					fmt.Sprintf(`SELECT .+ FROM %s WHERE name = .+ LIMIT 1;`,
						core.UserTable)).
					WithArgs(name).
					WillReturnRows(sqlmock.NewRows(core.UserRowNames))
			})

			It("returns nil *User", func() {
				Expect(user).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns ErrUserNotFound", func() {
				Expect(err).NotTo(Succeed())
				Expect(errors.Is(err, core.ErrUserNotFound)).To(BeTrue())
			})
		})

		When("no errors are met", func() {
			BeforeEach(func() {
				selectOK()
			})

			It("returns non-nil *User", func() {
				Expect(user).NotTo(BeNil())
			})
			It("succeeds", func() {
				Expect(err).To(Succeed())
			})
		})
	})

	Describe("UpdateUserCount", func() {
		var (
			ctx  context.Context
			user *core.User

			err error

			updateOK = func() {
				dbMock.ExpectExec(fmt.Sprintf(`UPDATE %s SET count = .+`,
					core.UserTable)).
					WithArgs(user.Count, user.Id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}
		)

		BeforeEach(func() {
			ctx = context.Background()
			user = &core.User{}
		})

		JustBeforeEach(func() {
			err = ur.UpdateUserCount(ctx, user)
		})

		When("user is nil", func() {
			BeforeEach(func() {
				user = nil
			})

			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("UPDATE query fails", func() {
			BeforeEach(func() {
				dbMock.ExpectExec(fmt.Sprintf(`UPDATE %s SET count = .+`,
					core.UserTable)).
					WithArgs(user.Count, user.Id).
					WillReturnError(errors.New(""))
			})
		})

		When("no errors are met", func() {
			BeforeEach(func() {
				updateOK()
			})

			It("succeeds", func() {
				Expect(err).To(Succeed())
			})
		})
	})
})

func mockUserRepository() (*UserRepository, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	Expect(db).NotTo(BeNil())
	Expect(mock).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	var ur *UserRepository
	Expect(func() { ur = NewUserRepository(db) }).NotTo(Panic())
	Expect(ur).NotTo(BeNil())

	return ur, mock
}

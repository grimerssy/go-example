package biz

import (
	"context"
	"errors"

	"github.com/grimerssy/go-example/internal/biz/mocks"
	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/consts"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewGreeterUseCase", func() {
	var (
		guc *GreeterUseCase

		userRepositoryMock *mocks.UserRepository

		test func()
	)

	JustBeforeEach(func() {
		test = func() {
			guc = NewGreeterUseCase(userRepositoryMock)
		}
	})

	When("userRepository is nil", func() {
		BeforeEach(func() {
			userRepositoryMock = nil
		})

		It("panics", func() {
			Expect(test).To(Panic())
		})
		It("returns nil", func() {
			Expect(guc).To(BeNil())
		})
	})

	When("userRepository is not nil", func() {
		BeforeEach(func() {
			userRepositoryMock = mocks.NewUserRepository(GinkgoT())
		})

		It("does not panic", func() {
			Expect(test).NotTo(Panic())
		})
		It("returns non-nil *GreeterUseCase", func() {
			Expect(guc).NotTo(BeNil())
		})
	})
})

var _ = Describe("GreeterUseCase", func() {
	var (
		guc *GreeterUseCase

		userRepositoryMock *mocks.UserRepository
	)

	BeforeEach(func() {
		userRepositoryMock = mocks.NewUserRepository(GinkgoT())
		guc = NewGreeterUseCase(userRepositoryMock)
	})

	Describe("Greet", func() {
		var (
			ctx    context.Context
			userId int64

			message string
			err     error

			user         *core.User
			expectedUser = &core.User{Count: 1}

			getUserByIdOK = func() {
				userRepositoryMock.EXPECT().
					GetUserById(ctx, userId).
					Return(user, nil)
			}
			updateUserCountOK = func() {
				userRepositoryMock.EXPECT().
					UpdateUserCount(ctx, user).
					Return(nil)
			}
		)

		BeforeEach(func() {
			ctx = context.Background()
			user = &core.User{}
		})

		JustBeforeEach(func() {
			message, err = guc.Greet(ctx, userId)
		})

		When("getting user by id fails", func() {
			BeforeEach(func() {
				userRepositoryMock.EXPECT().
					GetUserById(ctx, userId).
					Return(nil, errors.New(""))
			})

			It("returns empty string", func() {
				Expect(message).To(BeEmpty())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("getting user by id fails with ErrUserNotFound", func() {
			BeforeEach(func() {
				userRepositoryMock.EXPECT().
					GetUserById(ctx, userId).
					Return(nil, consts.ErrUserNotFound)
			})

			It("returns empty string", func() {
				Expect(message).To(BeEmpty())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns ErrUserNotFound", func() {
				Expect(errors.Is(err, consts.ErrUserNotFound)).To(BeTrue())
			})
		})

		When("updating user count fails", func() {
			BeforeEach(func() {
				getUserByIdOK()

				userRepositoryMock.EXPECT().
					UpdateUserCount(ctx, user).
					Return(errors.New(""))
			})

			It("returns empty string", func() {
				Expect(message).To(BeEmpty())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("no errors are met", func() {
			BeforeEach(func() {
				getUserByIdOK()
				updateUserCountOK()
			})

			It("returns non-empty string", func() {
				Expect(message).NotTo(BeEmpty())
			})
			It("succeeds", func() {
				Expect(err).To(Succeed())
			})
			It("matches user to an expected one", func() {
				Expect(user).To(Equal(expectedUser))
			})
		})
	})
})

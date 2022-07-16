package v1

import (
	"context"
	"errors"

	v1 "github.com/grimerssy/go-example/internal/api/v1"
	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/internal/service/v1/mocks"
	authMocks "github.com/grimerssy/go-example/pkg/auth/mocks"
	"github.com/grimerssy/go-example/pkg/consts"
	"github.com/grimerssy/go-example/pkg/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
)

var _ = Describe("NewAuthService", func() {
	var (
		as *AuthService

		authUseCaseMock *mocks.AuthUseCase

		test func()
	)

	JustBeforeEach(func() {
		test = func() {
			as = NewAuthService(authUseCaseMock)
		}
	})

	When("authUseCase is nil", func() {
		BeforeEach(func() {
			authUseCaseMock = nil
		})

		It("panics", func() {
			Expect(test).To(Panic())
		})
		It("returns nil", func() {
			Expect(as).To(BeNil())
		})
	})

	When("authUseCase is not nil", func() {
		BeforeEach(func() {
			authUseCaseMock = mocks.NewAuthUseCase(GinkgoT())
		})

		It("does not panic", func() {
			Expect(test).NotTo(Panic())
		})
		It("returns non-nil *AuthService", func() {
			Expect(as).NotTo(BeNil())
		})
	})
})

var _ = Describe("AuthService", func() {
	var (
		as *AuthService

		authUseCaseMock *mocks.AuthUseCase
	)

	BeforeEach(func() {
		authUseCaseMock = mocks.NewAuthUseCase(GinkgoT())
		as = NewAuthService(authUseCaseMock)
	})

	Describe("Signup", func() {
		var (
			ctx context.Context
			req *v1.SignupRequest

			res *v1.SignupResponse
			err error

			errCode          codes.Code
			errMsg           string
			errHasLogMessage bool
		)

		BeforeEach(func() {
			ctx = context.Background()
			req = &v1.SignupRequest{}
		})

		JustBeforeEach(func() {
			res, err = as.Signup(ctx, req)

			status := statusFromError(err)
			errCode = status.Code()
			errMsg = status.Message()
			_, errHasLogMessage = err.(log.Messenger)
		})

		When("ErrUserAlreadyExists occurs", func() {
			BeforeEach(func() {
				authUseCaseMock.EXPECT().
					Signup(context.Background(), &core.User{}).
					Return(consts.ErrUserAlreadyExists)
			})

			It("returns nil", func() {
				Expect(res).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns error with AlreadyExists code", func() {
				Expect(errCode).To(Equal(codes.AlreadyExists))
			})
			It("returns error with ErrUserAlreadyExists message", func() {
				Expect(errMsg).To(Equal(consts.ErrUserAlreadyExists.Error()))
			})
			It("returns error with log message", func() {
				Expect(errHasLogMessage).To(BeTrue())
			})
		})

		When("unexpected error occurs", func() {
			BeforeEach(func() {
				authUseCaseMock.EXPECT().
					Signup(context.Background(), &core.User{}).
					Return(errors.New(""))
			})

			It("returns nil", func() {
				Expect(res).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns error with Unknown code", func() {
				Expect(errCode).To(Equal(codes.Unknown))
			})
			It("returns error with UnexpectedErrorMessage", func() {
				Expect(errMsg).To(Equal(consts.UnexpectedErrorMessage))
			})
			It("returns error with log message", func() {
				Expect(errHasLogMessage).To(BeTrue())
			})
		})

		When("no errors are met", func() {
			BeforeEach(func() {
				authUseCaseMock.EXPECT().
					Signup(context.Background(), &core.User{}).
					Return(nil)
			})

			It("returns non-nil response", func() {
				Expect(res).NotTo(BeNil())
			})
			It("succeeds", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Login", func() {
		var (
			ctx context.Context
			req *v1.LoginRequest

			res *v1.LoginResponse
			err error

			errCode          codes.Code
			errMsg           string
			errHasLogMessage bool
		)

		BeforeEach(func() {
			ctx = context.Background()
			req = &v1.LoginRequest{}
		})

		JustBeforeEach(func() {
			res, err = as.Login(ctx, req)

			status := statusFromError(err)
			errCode = status.Code()
			errMsg = status.Message()
			_, errHasLogMessage = err.(log.Messenger)
		})

		When("ErrUserNotFound occurs", func() {
			BeforeEach(func() {
				authUseCaseMock.EXPECT().
					Login(context.Background(), &core.User{}).
					Return(nil, consts.ErrUserNotFound)
			})

			It("returns nil", func() {
				Expect(res).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns error with NotFound code", func() {
				Expect(errCode).To(Equal(codes.NotFound))
			})
			It("returns error with ErrUserNotFound message", func() {
				Expect(errMsg).To(Equal(consts.ErrUserNotFound.Error()))
			})
			It("returns error with log message", func() {
				Expect(errHasLogMessage).To(BeTrue())
			})
		})

		When("ErrInvalidPassword occurs", func() {
			BeforeEach(func() {
				authUseCaseMock.EXPECT().
					Login(context.Background(), &core.User{}).
					Return(nil, consts.ErrInvalidPassword)
			})

			It("returns nil", func() {
				Expect(res).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns error with Unauthenticated code", func() {
				Expect(errCode).To(Equal(codes.Unauthenticated))
			})
			It("returns error with ErrInvalidPassword message", func() {
				Expect(errMsg).To(Equal(consts.ErrInvalidPassword.Error()))
			})
			It("returns error with log message", func() {
				Expect(errHasLogMessage).To(BeTrue())
			})
		})

		When("unexpected error occurs", func() {
			BeforeEach(func() {
				authUseCaseMock.EXPECT().
					Login(context.Background(), &core.User{}).
					Return(nil, errors.New(""))
			})

			It("returns nil", func() {
				Expect(res).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns error with Unknown code", func() {
				Expect(errCode).To(Equal(codes.Unknown))
			})
			It("returns error with UnexpectedErrorMessage", func() {
				Expect(errMsg).To(Equal(consts.UnexpectedErrorMessage))
			})
			It("returns error with log message", func() {
				Expect(errHasLogMessage).To(BeTrue())
			})
		})

		When("no errors are met", func() {
			BeforeEach(func() {
				tokensMock := authMocks.NewTokens(GinkgoT())
				tokensMock.EXPECT().
					AccessToken().
					Return("")

				authUseCaseMock.EXPECT().
					Login(context.Background(), &core.User{}).
					Return(tokensMock, nil)
			})

			It("returns non-nil response", func() {
				Expect(res).NotTo(BeNil())
			})
			It("succeeds", func() {
				Expect(err).To(Succeed())
			})
		})
	})
})

package v1

import (
	"context"
	"errors"

	v1 "github.com/grimerssy/go-example/internal/api/v1"
	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/internal/service/v1/mocks"
	authMocks "github.com/grimerssy/go-example/pkg/auth/mocks"
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

			statusCode    codes.Code
			statusMessage string
		)

		BeforeEach(func() {
			ctx = context.Background()
			req = &v1.SignupRequest{}
		})

		JustBeforeEach(func() {
			res, err = as.Signup(ctx, req)

			status := statusFromError(err)
			statusCode = status.Code()
			statusMessage = status.Message()
		})

		When("errUserAlreadyExists occurs", func() {
			BeforeEach(func() {
				authUseCaseMock.EXPECT().
					Signup(context.Background(), &core.User{}).
					Return(errUserAlreadyExists)
			})

			It("returns nil", func() {
				Expect(res).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns error with AlreadyExists code", func() {
				Expect(statusCode).To(Equal(codes.AlreadyExists))
			})
			It("returns error with errUserAlreadyExists message", func() {
				Expect(statusMessage).To(Equal(errUserAlreadyExists.Error()))
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
				Expect(statusCode).To(Equal(codes.Unknown))
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

			statusCode    codes.Code
			statusMessage string
		)

		BeforeEach(func() {
			ctx = context.Background()
			req = &v1.LoginRequest{}
		})

		JustBeforeEach(func() {
			res, err = as.Login(ctx, req)

			status := statusFromError(err)
			statusCode = status.Code()
			statusMessage = status.Message()
		})

		When("errUserNotFound occurs", func() {
			BeforeEach(func() {
				authUseCaseMock.EXPECT().
					Login(context.Background(), &core.User{}).
					Return(nil, errUserNotFound)
			})

			It("returns nil", func() {
				Expect(res).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns error with NotFound code", func() {
				Expect(statusCode).To(Equal(codes.NotFound))
			})
			It("returns error with errUserNotFound message", func() {
				Expect(statusMessage).To(Equal(errUserNotFound.Error()))
			})
		})

		When("errInvalidPassword occurs", func() {
			BeforeEach(func() {
				authUseCaseMock.EXPECT().
					Login(context.Background(), &core.User{}).
					Return(nil, errInvalidPassword)
			})

			It("returns nil", func() {
				Expect(res).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns error with Unauthenticated code", func() {
				Expect(statusCode).To(Equal(codes.Unauthenticated))
			})
			It("returns error with errInvalidPassword message", func() {
				Expect(statusMessage).To(Equal(errInvalidPassword.Error()))
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
				Expect(statusCode).To(Equal(codes.Unknown))
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

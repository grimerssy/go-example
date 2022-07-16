package v1

import (
	"context"
	"errors"

	v1 "github.com/grimerssy/go-example/internal/api/v1"
	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/internal/service/v1/mocks"
	"github.com/grimerssy/go-example/pkg/consts"
	"github.com/grimerssy/go-example/pkg/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
)

var _ = Describe("NewGreeterService", func() {
	var (
		gs *GreeterService

		greeterUseCaseMock *mocks.GreeterUseCase

		test func()
	)

	JustBeforeEach(func() {
		test = func() {
			gs = NewGreeterService(greeterUseCaseMock)
		}
	})

	When("greeterUseCase is nil", func() {
		BeforeEach(func() {
			greeterUseCaseMock = nil
		})

		It("panics", func() {
			Expect(test).To(Panic())
		})
		It("returns nil", func() {
			Expect(gs).To(BeNil())
		})
	})

	When("greeterUseCaseMock is not nil", func() {
		BeforeEach(func() {
			greeterUseCaseMock = mocks.NewGreeterUseCase(GinkgoT())
		})

		It("does not panic", func() {
			Expect(test).NotTo(Panic())
		})
		It("returns non-nil *GreeterService", func() {
			Expect(gs).NotTo(BeNil())
		})
	})
})

var _ = Describe("GreeterService", func() {
	var (
		gs *GreeterService

		greeterUseCaseMock *mocks.GreeterUseCase
	)

	BeforeEach(func() {
		greeterUseCaseMock = mocks.NewGreeterUseCase(GinkgoT())
		gs = NewGreeterService(greeterUseCaseMock)
	})

	Describe("Greet", func() {
		var (
			userId int64

			ctx context.Context
			req *v1.GreetRequest

			res *v1.GreetResponse
			err error

			errCode          codes.Code
			errMsg           string
			errHasLogMessage bool
		)

		BeforeEach(func() {
			ctx = context.WithValue(context.Background(), core.UserIdKey, userId)
			req = &v1.GreetRequest{}
		})

		JustBeforeEach(func() {
			res, err = gs.Greet(ctx, req)

			status := statusFromError(err)
			errCode = status.Code()
			errMsg = status.Message()
			_, errHasLogMessage = err.(log.Messenger)
		})

		When("ErrContextHasNoUserId occurs", func() {
			BeforeEach(func() {
				greeterUseCaseMock.EXPECT().
					Greet(ctx, userId).
					Return("", consts.ErrContextHasNoUserId)
			})

			It("returns nil", func() {
				Expect(res).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns error with Internal code", func() {
				Expect(errCode).To(Equal(codes.Internal))
			})
			It("returns error with ErrContextHasNoUserId message", func() {
				Expect(errMsg).To(Equal(consts.ErrContextHasNoUserId.Error()))
			})
			It("returns error with log message", func() {
				Expect(errHasLogMessage).To(BeTrue())
			})
		})

		When("ErrUserNotFound occurs", func() {
			BeforeEach(func() {
				greeterUseCaseMock.EXPECT().
					Greet(ctx, userId).
					Return("", consts.ErrUserNotFound)
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

		When("unexpected error occurs", func() {
			BeforeEach(func() {
				greeterUseCaseMock.EXPECT().
					Greet(ctx, userId).
					Return("", errors.New(""))
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
				greeterUseCaseMock.EXPECT().
					Greet(ctx, userId).
					Return("", nil)
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

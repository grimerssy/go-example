package v1

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	v1 "github.com/grimerssy/go-example/internal/api/v1"
	"github.com/grimerssy/go-example/internal/core"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
)

var _ = Describe("NewGreeterService", func() {
	var (
		gs *GreeterService

		ctrl               *gomock.Controller
		greeterUseCaseMock *greeterUseCaseMock

		test func()
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

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
			greeterUseCaseMock = NewgreeterUseCaseMock(ctrl)
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

		ctrl               *gomock.Controller
		greeterUseCaseMock *greeterUseCaseMock
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		greeterUseCaseMock = NewgreeterUseCaseMock(ctrl)
		gs = NewGreeterService(greeterUseCaseMock)
	})

	Describe("Greet", func() {
		var (
			userId int64

			ctx context.Context
			req *v1.GreetRequest

			res *v1.GreetResponse
			err error

			statusCode    codes.Code
			statusMessage string
		)

		BeforeEach(func() {
			ctx = context.WithValue(context.Background(), core.UserIdKey, userId)
			req = &v1.GreetRequest{}
		})

		JustBeforeEach(func() {
			res, err = gs.Greet(ctx, req)

			status := statusFromError(err)
			statusCode = status.Code()
			statusMessage = status.Message()
		})

		When("errContextHasNoUserId occurs", func() {
			BeforeEach(func() {
				greeterUseCaseMock.EXPECT().
					Greet(ctx, userId).
					Return("", errContextHasNoUserId)
			})

			It("returns nil", func() {
				Expect(res).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns error with Internal code", func() {
				Expect(statusCode).To(Equal(codes.Internal))
			})
			It("returns error with errContextHasNoUserId message", func() {
				Expect(statusMessage).To(Equal(errContextHasNoUserId.Error()))
			})
		})

		When("errUserNotFound occurs", func() {
			BeforeEach(func() {
				greeterUseCaseMock.EXPECT().
					Greet(ctx, userId).
					Return("", errUserNotFound)
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
				Expect(statusCode).To(Equal(codes.Unknown))
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

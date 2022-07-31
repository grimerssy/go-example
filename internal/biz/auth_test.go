package biz

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/auth"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewAuthUseCase", func() {
	var (
		auc *AuthUseCase

		ctrl               *gomock.Controller
		tokenManagerMock   *tokenManagerMock
		idObfuscatorMock   *idObfuscatorMock
		passwordHasherMock *passwordHasherMock
		userRepositoryMock *userRepositoryMock

		test func()
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

	JustBeforeEach(func() {
		test = func() {
			auc = NewAuthUseCase(
				tokenManagerMock,
				idObfuscatorMock,
				passwordHasherMock,
				userRepositoryMock,
			)
		}
	})

	When("tokenManager is nil", func() {
		BeforeEach(func() {
			tokenManagerMock = nil
			idObfuscatorMock = NewidObfuscatorMock(ctrl)
			passwordHasherMock = NewpasswordHasherMock(ctrl)
			userRepositoryMock = NewuserRepositoryMock(ctrl)
		})

		It("panics", func() {
			Expect(test).To(Panic())
		})
		It("returns nil", func() {
			Expect(auc).To(BeNil())
		})
	})

	When("idObfuscator is nil", func() {
		BeforeEach(func() {
			tokenManagerMock = NewtokenManagerMock(ctrl)
			idObfuscatorMock = nil
			passwordHasherMock = NewpasswordHasherMock(ctrl)
			userRepositoryMock = NewuserRepositoryMock(ctrl)
		})

		It("panics", func() {
			Expect(test).To(Panic())
		})
		It("returns nil", func() {
			Expect(auc).To(BeNil())
		})
	})

	When("passwordHasher is nil", func() {
		BeforeEach(func() {
			tokenManagerMock = NewtokenManagerMock(ctrl)
			idObfuscatorMock = NewidObfuscatorMock(ctrl)
			passwordHasherMock = nil
			userRepositoryMock = NewuserRepositoryMock(ctrl)
		})

		It("panics", func() {
			Expect(test).To(Panic())
		})
		It("returns nil", func() {
			Expect(auc).To(BeNil())
		})
	})

	When("userRepository is nil", func() {
		BeforeEach(func() {
			tokenManagerMock = NewtokenManagerMock(ctrl)
			idObfuscatorMock = NewidObfuscatorMock(ctrl)
			passwordHasherMock = NewpasswordHasherMock(ctrl)
			userRepositoryMock = nil
		})

		It("panics", func() {
			Expect(test).To(Panic())
		})
		It("returns nil", func() {
			Expect(auc).To(BeNil())
		})
	})

	When("none of the parameters are nil", func() {
		BeforeEach(func() {
			tokenManagerMock = NewtokenManagerMock(ctrl)
			idObfuscatorMock = NewidObfuscatorMock(ctrl)
			passwordHasherMock = NewpasswordHasherMock(ctrl)
			userRepositoryMock = NewuserRepositoryMock(ctrl)
		})

		It("does not panic", func() {
			Expect(test).NotTo(Panic())
		})
		It("returns non-nil *AuthUseCase", func() {
			Expect(auc).NotTo(BeNil())
		})
	})
})

var _ = Describe("AuthUseCase", func() {
	var (
		auc *AuthUseCase

		ctrl               *gomock.Controller
		tokenManagerMock   *tokenManagerMock
		idObfuscatorMock   *idObfuscatorMock
		passwordHasherMock *passwordHasherMock
		userRepositoryMock *userRepositoryMock
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		idObfuscatorMock = NewidObfuscatorMock(ctrl)
		tokenManagerMock = NewtokenManagerMock(ctrl)
		passwordHasherMock = NewpasswordHasherMock(ctrl)
		userRepositoryMock = NewuserRepositoryMock(ctrl)
		auc = NewAuthUseCase(
			tokenManagerMock,
			idObfuscatorMock,
			passwordHasherMock,
			userRepositoryMock,
		)
	})

	Describe("Signup", func() {
		var (
			ctx  context.Context
			user *core.User

			err error

			hashPasswordOK = func() {
				passwordHasherMock.EXPECT().
					HashPassword(user.Password).
					Return(user.Password, nil)
			}
			createUserOK = func() {
				userRepositoryMock.EXPECT().
					CreateUser(ctx, user).
					Return(nil)
			}
		)

		BeforeEach(func() {
			ctx = context.Background()
			user = &core.User{}
		})

		JustBeforeEach(func() {
			err = auc.Signup(ctx, user)
		})

		When("password hashing fails", func() {
			BeforeEach(func() {
				passwordHasherMock.EXPECT().
					HashPassword(user.Password).
					Return("", errors.New(""))
			})

			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("creating user fails", func() {
			BeforeEach(func() {
				hashPasswordOK()

				userRepositoryMock.EXPECT().
					CreateUser(ctx, user).
					Return(errors.New(""))
			})

			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("creating user fails with errUserAlreadyExists", func() {
			BeforeEach(func() {
				hashPasswordOK()

				userRepositoryMock.EXPECT().
					CreateUser(ctx, user).
					Return(errUserAlreadyExists)
			})

			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns errUserAlreadyExists", func() {
				Expect(errors.Is(err, errUserAlreadyExists)).To(BeTrue())
			})
		})

		When("no errors are met", func() {
			BeforeEach(func() {
				hashPasswordOK()
				createUserOK()
			})

			It("succeeds", func() {
				Expect(err).To(Succeed())
			})
		})
	})

	Describe("Login", func() {
		var (
			ctx  context.Context
			user *core.User

			token auth.Token
			err   error

			getUserByNameOK = func() {
				userRepositoryMock.EXPECT().
					GetUserByName(ctx, user.Name).
					Return(user, nil)
			}
			isPasswordEqualToHashOK = func() {
				passwordHasherMock.EXPECT().
					IsPasswordEqualToHash(user.Password, user.Password).
					Return(true)
			}
			obfuscateIdOK = func() {
				idObfuscatorMock.EXPECT().
					ObfuscateId(user.Id).
					Return(user.Id, nil)
			}
			generateTokensOK = func() {
				claims := map[string]string{
					core.UserIdKey: fmt.Sprintf("%v", user.Id),
				}
				tokenManagerMock.EXPECT().
					GenerateToken(claims).
					Return(auth.NewtokenMock(ctrl), nil)
			}
		)

		BeforeEach(func() {
			ctx = context.Background()
			user = &core.User{}
		})

		JustBeforeEach(func() {
			token, err = auc.Login(ctx, user)
		})

		When("getting user by name fails", func() {
			BeforeEach(func() {
				userRepositoryMock.EXPECT().
					GetUserByName(ctx, user.Name).
					Return(nil, errors.New(""))
			})

			It("returns nil", func() {
				Expect(token).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("getting user by name fails with errUserNotFound", func() {
			BeforeEach(func() {
				userRepositoryMock.EXPECT().
					GetUserByName(ctx, user.Name).
					Return(nil, errUserNotFound)
			})

			It("returns nil", func() {
				Expect(token).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns errUserNotFound", func() {
				Expect(errors.Is(err, errUserNotFound)).To(BeTrue())
			})
		})

		When("password checking fails", func() {
			BeforeEach(func() {
				getUserByNameOK()

				passwordHasherMock.EXPECT().
					IsPasswordEqualToHash(user.Password, user.Password).
					Return(false)
			})

			It("returns nil", func() {
				Expect(token).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns errInvalidPassword", func() {
				Expect(errors.Is(err, errInvalidPassword)).To(BeTrue())
			})
		})

		When("id obfuscation fails", func() {
			BeforeEach(func() {
				getUserByNameOK()
				isPasswordEqualToHashOK()

				idObfuscatorMock.EXPECT().
					ObfuscateId(user.Id).
					Return(int64(0), errors.New(""))
			})

			It("returns nil", func() {
				Expect(token).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("token generation fails", func() {
			BeforeEach(func() {
				getUserByNameOK()
				isPasswordEqualToHashOK()
				obfuscateIdOK()

				tokenManagerMock.EXPECT().
					GenerateToken(map[string]string{
						core.UserIdKey: fmt.Sprintf("%v", user.Id),
					}).
					Return(nil, errors.New(""))
			})

			It("returns nil", func() {
				Expect(token).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("no errors are met", func() {
			BeforeEach(func() {
				getUserByNameOK()
				isPasswordEqualToHashOK()
				obfuscateIdOK()
				generateTokensOK()
			})

			It("returns non-nil token", func() {
				Expect(token).NotTo(BeNil())
			})
			It("succeeds", func() {
				Expect(err).To(Succeed())
			})
		})
	})

	Describe("GetUserId", func() {
		var (
			ctx   context.Context
			token auth.AccessToken

			userId int64
			err    error

			parseTokensOK = func() {
				claims := map[string]string{
					core.UserIdKey: fmt.Sprintf("%v", userId),
				}
				tokenManagerMock.EXPECT().
					ParseToken(token).
					Return(claims, nil)
			}
			deobfuscateIdOK = func() {
				idObfuscatorMock.EXPECT().
					DeobfuscateId(userId).
					Return(userId, nil)
			}
		)

		BeforeEach(func() {
			ctx = context.Background()
			token = auth.NewaccessTokenMock(ctrl)
		})

		JustBeforeEach(func() {
			userId, err = auc.GetUserId(ctx, token)
		})

		When("token parsing fails", func() {
			BeforeEach(func() {
				tokenManagerMock.EXPECT().
					ParseToken(token).
					Return(nil, errors.New(""))
			})

			It("returns zero", func() {
				Expect(userId).To(BeZero())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("id deobfuscation fails", func() {
			BeforeEach(func() {
				parseTokensOK()

				idObfuscatorMock.EXPECT().
					DeobfuscateId(userId).
					Return(int64(0), errors.New(""))
			})

			It("returns zero", func() {
				Expect(userId).To(BeZero())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("no errors are met", func() {
			BeforeEach(func() {
				parseTokensOK()
				deobfuscateIdOK()
			})

			It("succeeds", func() {
				Expect(err).To(Succeed())
			})
		})
	})
})

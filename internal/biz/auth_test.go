package biz

import (
	"context"
	"errors"

	"github.com/grimerssy/go-example/internal/biz/mocks"
	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/auth"
	authMock "github.com/grimerssy/go-example/pkg/auth/mocks"
	"github.com/grimerssy/go-example/pkg/consts"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewAuthUseCase", func() {
	var (
		auc *AuthUseCase

		tokenManagerMock   *mocks.TokenManager
		idObfuscatorMock   *mocks.IdObfuscator
		passwordHasherMock *mocks.PasswordHasher
		userRepositoryMock *mocks.UserRepository

		test func()
	)

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
			idObfuscatorMock = mocks.NewIdObfuscator(GinkgoT())
			passwordHasherMock = mocks.NewPasswordHasher(GinkgoT())
			userRepositoryMock = mocks.NewUserRepository(GinkgoT())
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
			tokenManagerMock = mocks.NewTokenManager(GinkgoT())
			idObfuscatorMock = nil
			passwordHasherMock = mocks.NewPasswordHasher(GinkgoT())
			userRepositoryMock = mocks.NewUserRepository(GinkgoT())
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
			tokenManagerMock = mocks.NewTokenManager(GinkgoT())
			idObfuscatorMock = mocks.NewIdObfuscator(GinkgoT())
			passwordHasherMock = nil
			userRepositoryMock = mocks.NewUserRepository(GinkgoT())
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
			tokenManagerMock = mocks.NewTokenManager(GinkgoT())
			idObfuscatorMock = mocks.NewIdObfuscator(GinkgoT())
			passwordHasherMock = mocks.NewPasswordHasher(GinkgoT())
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
			tokenManagerMock = mocks.NewTokenManager(GinkgoT())
			idObfuscatorMock = mocks.NewIdObfuscator(GinkgoT())
			passwordHasherMock = mocks.NewPasswordHasher(GinkgoT())
			userRepositoryMock = mocks.NewUserRepository(GinkgoT())
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

		tokenManagerMock   *mocks.TokenManager
		idObfuscatorMock   *mocks.IdObfuscator
		passwordHasherMock *mocks.PasswordHasher
		userRepositoryMock *mocks.UserRepository
	)

	BeforeEach(func() {
		idObfuscatorMock = mocks.NewIdObfuscator(GinkgoT())
		tokenManagerMock = mocks.NewTokenManager(GinkgoT())
		passwordHasherMock = mocks.NewPasswordHasher(GinkgoT())
		userRepositoryMock = mocks.NewUserRepository(GinkgoT())
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

		When("creating user fails with ErrUserAlreadyExists", func() {
			BeforeEach(func() {
				hashPasswordOK()

				userRepositoryMock.EXPECT().
					CreateUser(ctx, user).
					Return(consts.ErrUserAlreadyExists)
			})

			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns ErrUserAlreadyExists", func() {
				Expect(errors.Is(err, consts.ErrUserAlreadyExists)).To(BeTrue())
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

			tokens auth.Tokens
			err    error

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
			defaultClaimsCALL = func() {
				tokenManagerMock.EXPECT().
					DefaultClaims().
					Return(authMock.NewClaims(GinkgoT()))
			}
			generateTokensOK = func() {
				tokenManagerMock.EXPECT().
					GenerateTokens(newUserIdClaims(tokenManagerMock.
						DefaultClaims(), user.Id)).
					Return(authMock.NewTokens(GinkgoT()), nil)
			}
		)

		BeforeEach(func() {
			ctx = context.Background()
			user = &core.User{}
		})

		JustBeforeEach(func() {
			tokens, err = auc.Login(ctx, user)
		})

		When("getting user by name fails", func() {
			BeforeEach(func() {
				userRepositoryMock.EXPECT().
					GetUserByName(ctx, user.Name).
					Return(nil, errors.New(""))
			})

			It("returns nil", func() {
				Expect(tokens).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
		})

		When("getting user by name fails with ErrUserNotFound", func() {
			BeforeEach(func() {
				userRepositoryMock.EXPECT().
					GetUserByName(ctx, user.Name).
					Return(nil, consts.ErrUserNotFound)
			})

			It("returns nil", func() {
				Expect(tokens).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns ErrUserNotFound", func() {
				Expect(errors.Is(err, consts.ErrUserNotFound)).To(BeTrue())
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
				Expect(tokens).To(BeNil())
			})
			It("fails", func() {
				Expect(err).NotTo(Succeed())
			})
			It("returns ErrInvalidPassword", func() {
				Expect(errors.Is(err, consts.ErrInvalidPassword)).To(BeTrue())
			})
		})

		When("id obfuscation fails", func() {
			BeforeEach(func() {
				getUserByNameOK()
				isPasswordEqualToHashOK()

				idObfuscatorMock.EXPECT().
					ObfuscateId(user.Id).
					Return(0, errors.New(""))
			})

			It("returns nil", func() {
				Expect(tokens).To(BeNil())
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
				defaultClaimsCALL()

				tokenManagerMock.EXPECT().
					GenerateTokens(newUserIdClaims(tokenManagerMock.
						DefaultClaims(), user.Id)).
					Return(nil, errors.New(""))
			})

			It("returns nil", func() {
				Expect(tokens).To(BeNil())
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
				defaultClaimsCALL()
				generateTokensOK()
			})

			It("returns non-nil tokens", func() {
				Expect(tokens).NotTo(BeNil())
			})
			It("succeeds", func() {
				Expect(err).To(Succeed())
			})
		})
	})

	Describe("GetUserId", func() {
		var (
			ctx    context.Context
			tokens auth.Tokens

			userId int64
			err    error

			parseTokensOK = func() {
				claimsBase := authMock.NewClaims(GinkgoT())
				claims := newUserIdClaims(claimsBase, userId)
				tokenManagerMock.EXPECT().
					ParseTokens(tokens, &userIdClaims{}).
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
			tokens = authMock.NewTokens(GinkgoT())
		})

		JustBeforeEach(func() {
			userId, err = auc.GetUserId(ctx, tokens)
		})

		When("tokens parsing fails", func() {
			BeforeEach(func() {
				tokenManagerMock.EXPECT().
					ParseTokens(tokens, &userIdClaims{}).
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
					Return(0, errors.New(""))
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

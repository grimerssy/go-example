package auth

//go:generate mockgen -source=tokens.go -destination=tokens_mock.go -package=auth -mock_names=Token=tokenMock,AccessToken=accessTokenMock
type Token interface {
	AccessToken
}

//go:generate mockery --name=AccessToken --with-expecter --quiet
type AccessToken interface {
	AccessToken() string
}

type tokens struct {
	access AccessToken
}

func NewTokens(access string) *tokens {
	return &tokens{
		access: NewAccessToken(access),
	}
}

func (t *tokens) AccessToken() string {
	return t.access.AccessToken()
}

type accessToken struct {
	token string
}

func NewAccessToken(token string) *accessToken {
	return &accessToken{
		token: token,
	}
}

func (a *accessToken) AccessToken() string {
	return a.token
}

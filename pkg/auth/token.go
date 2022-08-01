package auth

//go:generate mockgen -source=token.go -destination=token_mock.go -package=auth -mock_names=Token=tokenMock,AccessToken=accessTokenMock
type Token interface {
	AccessToken
}

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

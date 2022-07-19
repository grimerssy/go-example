package auth

//go:generate mockery --name=Tokens --with-expecter
type Tokens interface {
	AccessToken
}

//go:generate mockery --name=AccessToken --with-expecter
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

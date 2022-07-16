package auth

//go:generate mockery --name=Tokens --with-expecter
type Tokens interface {
	accessToken
}

type accessToken interface {
	AccessToken() string
}

type tokens struct {
	accessToken string
}

func newTokens(accessToken string) *tokens {
	return &tokens{
		accessToken: accessToken,
	}
}

func (t *tokens) AccessToken() string {
	return t.accessToken
}

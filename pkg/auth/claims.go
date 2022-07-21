package auth

//go:generate mockery --name=Claims --with-expecter --quiet
type Claims interface {
	Valid() error
}

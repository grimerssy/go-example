package auth

//go:generate mockery --name=Claims --with-expecter
type Claims interface {
	Valid() error
}

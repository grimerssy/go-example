package validation

//go:generate mockery --name=Validator --with-expecter --quiet
type Validator interface {
	Validate() error
}

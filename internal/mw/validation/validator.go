package validation

type validator interface {
	Validate() error
}

package validation

//go:generate mockgen -source=validator.go -destination=validator_mock.go -package=validation -mock_names=Validator=validatorMock
type Validator interface {
	Validate() error
}

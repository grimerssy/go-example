package consts

const (
	UnexpectedErrorMessage = "an unexpected error occurred"

	ErrInvalidPassword   = constError("invalid password")
	ErrUserAlreadyExists = constError("user with this name already exists")
	ErrUserNotFound      = constError("user was not found")

	ErrContextHasNoUserId = constError("context does not have user id")
)

type constError string

func (e constError) Error() string {
	return string(e)
}

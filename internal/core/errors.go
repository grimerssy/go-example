package core

const (
	ErrInvalidPassword   = constError("invalid password")
	ErrUserAlreadyExists = constError("user with this name already exists")
	ErrUserNotFound      = constError("user was not found")
)

type constError string

func (e constError) Error() string {
	return string(e)
}

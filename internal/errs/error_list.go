package errs

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrNotFound           = errors.New("not found")
	ErrInvalidUserID      = errors.New("invalid user ID")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidFieldValue  = errors.New("invalid field value")
	ErrInvalidUserName    = errors.New("invalid username")
	ErrNegativeID         = errors.New("negative id err. id should positive")
	ErrInvalidIDFormat    = errors.New("invalid user id format")
)

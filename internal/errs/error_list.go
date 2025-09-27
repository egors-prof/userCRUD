package errs

import "errors"


var (
	ErrUserNotFound       = errors.New("user not found")
	ErrNotFound           = errors.New("not found")
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidFieldValue  = errors.New("invalid field value")
	ErrInvalidUserName    = errors.New("invalid username")
	
)

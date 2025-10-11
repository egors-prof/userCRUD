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


var(
	ErrUserAlreadyExists=errors.New("user already exists")
	ErrBindJson=errors.New("error binding json")
	ErrCreatingUser=errors.New("error creating user")
	ErrHashing=errors.New("error hashing")
	ErrIncorrectUsernameOrPassword=errors.New("incorrect username or password")
	ErrInvalidToken=errors.New("invalid token")
	ErrSomethingWentWrong=errors.New("something went wrong")
	ErrAccessDenied=errors.New("access denied")
)
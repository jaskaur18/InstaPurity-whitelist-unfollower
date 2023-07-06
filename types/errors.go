package types

import "errors"

// LoginError Login Errors
type LoginError struct {
	Msg string
}

func (e LoginError) Error() string {
	return e.Msg
}

var (
	ErrUsernamePassword = errors.New("username or password are empty")
	ErrLoginImport      = errors.New("login error: import session failed")
	ErrLogin2FARequired = errors.New("login error: 2FA required")
)

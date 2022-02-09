package entity

import "errors"

//ErrCreateSession create session
const (
	ErrCreateSession = "fail on create session"
	ErrFindUser = "fail on find user"
)

//ErrNotFound not found
var ErrNotFound = errors.New("not found")

//ErrGeneratePassword generate password
var ErrGeneratePassword = errors.New("error on generate hash of password")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")
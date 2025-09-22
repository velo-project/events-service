package errors

import "errors"

var (
	ErrUserAlreadySubscribed = errors.New("Este usuário já está cadastrado nesse evento")
)

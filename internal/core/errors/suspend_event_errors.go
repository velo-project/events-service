package errors

import "errors"

var (
	ErrBlockedSuspendEvent  = errors.New("Não é possível suspender um evento que acontecerá em 3 dias ou menos")
	ErrBlockedActivateEvent = errors.New("Não é possível ativar um evento que acontecerá em 3 dias ou menos")
)

package errors

import "errors"

var (
	ErrBlockedCancelSubscription = errors.New("Não é possível cancelar a inscrição em um evento que acontecerá em 7 dias ou menos")
	ErrUserSubscriptionNotFound  = errors.New("Inscrição não encontrada para este evento e usuário")
	ErrEventNotFound             = errors.New("Esse evento não existe")
)

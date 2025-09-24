package ports

type ConfirmSubscriptionPort interface {
	Execute(code string, userId int, eventId int) error
}

package ports

type CancelSubscriptionPort interface {
	Execute(eventId int, userId int) error
}

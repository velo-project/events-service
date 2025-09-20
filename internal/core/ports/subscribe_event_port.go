package ports

type SubscribeEventPort interface {
	Execute(userId int, eventId int) error
}

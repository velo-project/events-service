package ports

type CancelEventPort interface {
	Execute(eventId int) error
}

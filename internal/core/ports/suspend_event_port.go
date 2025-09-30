package ports

type SuspendEventPort interface {
	Execute(eventId int) error
}

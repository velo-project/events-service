package ports

type ActivateEventPort interface {
	Execute(eventId int) error
}

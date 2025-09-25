package ports

import "time"

type GetConfirmationCodePort interface {
	Execute(userId int, eventId int) (*string, *time.Time, error)
}

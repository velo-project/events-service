package ports

import "io"

type SaveFilePort interface {
	Execute(file io.Reader) (*string, error)
}

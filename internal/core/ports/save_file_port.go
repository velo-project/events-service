package ports

import "io"

type SaveFilePort interface {
	Execute(file io.Reader, fileExtension string) (*string, error)
}

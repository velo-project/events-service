package storage

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
	"golang.org/x/net/context"
)

type saveFileAdapter struct {
	bucket *storage.BucketHandle
}

func NewSaveFileAdapter(bucket *storage.BucketHandle) ports.SaveFilePort {
	return &saveFileAdapter{
		bucket: bucket,
	}
}

func (s saveFileAdapter) Execute(file io.Reader) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	compressedImage, err := compressImage(file)
	if err != nil {
		return nil, err
	}

	fileNameBrute, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	fileName := fileNameBrute.String()

	now := time.Now()
	year, month, day := now.Date()

	completeObjectName := fmt.Sprintf("images/%d/%02d/%02d/users/%s%s",
		year,
		int(month),
		day,
		fileName,
		".jpeg",
	)

	object := s.bucket.Object(completeObjectName)
	wc := object.NewWriter(ctx)

	if _, err := io.Copy(wc, compressedImage); err != nil {
		wc.Close()
		return nil, fmt.Errorf("falha ao copiar o arquivo para o GCS: %w", err)
	}

	if err := wc.Close(); err != nil {
		return nil, fmt.Errorf("falha ao finalizar o upload no GCS: %w", err)
	}

	return &completeObjectName, nil
}

func compressImage(file io.Reader) (io.Reader, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	resizedImg := resize.Resize(1000, 0, img, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, resizedImg, nil)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

package filestorages

import (
	"context"

	"github.com/Lornzo/filestorages/localstorage"
)

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		Storage: &localstorage.LocalStorage{},
	}
}

type LocalStorage struct {
	Storage interface {
		Upload(ctx context.Context, data []byte) error
		SetMineType(minetype string) *localstorage.LocalStorage
		SetExtension(extension string) *localstorage.LocalStorage
		Object(object string) *localstorage.LocalStorage
		SetDirectory(dirs ...string) *localstorage.LocalStorage
		Bucket(bucket string) *localstorage.LocalStorage
	}
}

// Object interface
func (l *LocalStorage) NewWriter(ctx context.Context) Writer {
	return nil
}

func (l *LocalStorage) SetMineType(minetype string) Object {
	l.Storage.SetMineType(minetype)
	return l
}

func (l *LocalStorage) SetExtension(extension string) Object {
	l.Storage.SetExtension(extension)
	return l
}

func (l *LocalStorage) Upload(ctx context.Context, data []byte) error {
	return l.Storage.Upload(ctx, data)
}

// Bucket interface
func (l *LocalStorage) SetDirectory(dirs ...string) Bucket {
	l.Storage.SetDirectory(dirs...)
	return l
}

func (l *LocalStorage) Object(object string) Object {
	l.Storage.Object(object)
	return l
}

// Storage interface
func (l *LocalStorage) Bucket(bucket string) Bucket {
	l.Storage.Bucket(bucket)
	return l
}

func (l *LocalStorage) Close() error {
	return nil
}

package gcpstorage

import (
	"bytes"
	"context"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

func NewObject() *Object {
	return &Object{
		ObjectHandle: &storage.ObjectHandle{},
	}
}

type Object struct {
	*storage.ObjectHandle
}

func (o *Object) NewWriter(ctx context.Context) *Writer {
	return &Writer{Writer: o.ObjectHandle.NewWriter(ctx)}
}

func (o *Object) SetMineType(minetype string) *Object {
	return o
}

func (o *Object) SetExtension(extension string) *Object {
	return o
}

func (o *Object) Upload(ctx context.Context, data []byte) error {

	type iWriter interface {
		io.Writer
		SetChunkSize(size int)
		Close() error
	}

	var (
		writer iWriter
		cancel context.CancelFunc
		err    error
	)

	ctx, cancel = context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	writer = o.NewWriter(ctx)
	writer.SetChunkSize(0)

	if _, err = io.Copy(writer, bytes.NewBuffer(data)); err != nil {
		return err
	}

	if err = writer.Close(); err != nil {
		return err
	}
	return nil

}

func (o *Object) Delete(ctx context.Context) error {
	return o.ObjectHandle.Delete(ctx)
}

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

func (g *Object) SetMineType(minetype string) *Object {
	return g
}

func (g *Object) SetExtension(extension string) *Object {
	return g
}

func (g *Object) Upload(ctx context.Context, data []byte) error {

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

	writer = g.NewWriter(ctx)
	writer.SetChunkSize(0)

	if _, err = io.Copy(writer, bytes.NewBuffer(data)); err != nil {
		return err
	}

	if err = writer.Close(); err != nil {
		return err
	}
	return nil

}

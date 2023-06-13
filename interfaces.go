package filestorages

import (
	"context"
	"io"
)

type Writer interface {
	io.Writer
	SetChunkSize(size int)
	Close() error
}

type Storage interface {
	Bucket(bucket string) Bucket
	Close() error
}

type Object interface {
	NewWriter(ctx context.Context) Writer // has abendoned
	SetMineType(minetype string) Object
	SetExtension(extension string) Object
	Upload(ctx context.Context, data []byte) error
}

type Bucket interface {
	SetDirectory(dirs ...string) Bucket
	Object(object string) Object
}

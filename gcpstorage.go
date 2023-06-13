package filestorages

import (
	"context"

	"github.com/Lornzo/filestorages/gcpstorage"
)

type gcpObject struct {
	obj interface {
		NewWriter(ctx context.Context) *gcpstorage.Writer
		SetMineType(minetype string) *gcpstorage.Object
		SetExtension(extension string) *gcpstorage.Object
		Upload(ctx context.Context, data []byte) error
	}
}

func (g *gcpObject) NewWriter(ctx context.Context) Writer {
	return g.obj.NewWriter(ctx)
}

func (g *gcpObject) SetMineType(minetype string) Object {
	g.obj.SetMineType(minetype)
	return g
}

func (g *gcpObject) SetExtension(extension string) Object {
	g.obj.SetExtension(extension)
	return g
}

func (l *gcpObject) Upload(ctx context.Context, data []byte) error {
	return l.obj.Upload(ctx, data)
}

type gcpBucket struct {
	bucket interface {
		SetDirectory(dirs ...string) *gcpstorage.Bucket
		Object(object string) *gcpstorage.Object
	}
}

func (g *gcpBucket) SetDirectory(dirs ...string) Bucket {
	g.bucket.SetDirectory(dirs...)
	return g
}

func (g *gcpBucket) Object(object string) Object {
	return &gcpObject{
		obj: g.bucket.Object(object),
	}
}

func NewGCPStorage(ctx context.Context) (*GCPStorage, error) {

	var (
		theStorage *gcpstorage.Storage
		err        error
	)

	if theStorage, err = gcpstorage.NewStorage(ctx); err != nil {
		return nil, err
	}

	return &GCPStorage{
		Storage: theStorage,
	}, nil

}

type GCPStorage struct {
	Storage interface {
		Bucket(bucket string) *gcpstorage.Bucket
		Close() error
	}
}

func (g *GCPStorage) Bucket(bucket string) Bucket {
	return &gcpBucket{
		bucket: g.Storage.Bucket(bucket),
	}
}

func (g *GCPStorage) Close() error {
	return g.Storage.Close()
}

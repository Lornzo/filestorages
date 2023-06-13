package gcpstorage

import (
	"context"

	"cloud.google.com/go/storage"
)

func NewStorage(ctx context.Context) (*Storage, error) {
	var (
		client *storage.Client
		err    error
	)

	if client, err = storage.NewClient(ctx); err != nil {
		return nil, err
	}

	return &Storage{Client: client}, nil
}

type Storage struct {
	*storage.Client
}

func (g *Storage) Bucket(bucket string) *Bucket {

	var theBucket Bucket = Bucket{
		BucketHandle: g.Client.Bucket(bucket),
	}

	return &theBucket
}

func (g *Storage) Close() error {
	return g.Client.Close()
}

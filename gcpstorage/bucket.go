package gcpstorage

import (
	"fmt"
	"strings"

	"cloud.google.com/go/storage"
)

type Bucket struct {
	*storage.BucketHandle
	Directories []string
}

func (g *Bucket) SetDirectory(dirs ...string) *Bucket {
	g.Directories = dirs
	return g
}

func (g *Bucket) Object(object string) *Object {

	var name string = object
	if len(g.Directories) > 0 {
		name = fmt.Sprint(strings.Join(g.Directories, "/"), "/", object)
	}

	return &Object{
		ObjectHandle: g.BucketHandle.Object(name),
	}
}

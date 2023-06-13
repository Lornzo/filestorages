package localstorage

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
)

type LocalStorage struct {
	bucket      string
	object      string
	mineType    string
	extension   string
	directories []string
}

func (l *LocalStorage) Upload(ctx context.Context, data []byte) error {

	var (
		dir          string = l.bucket + "/" + strings.Join(l.directories, "/")
		err          error
		extension    string = l.getMineExtension(l.mineType)
		filePathName string
	)

	if extension == "" {
		return errors.New("not support mine type")
	}

	if err = l.mkdir(dir); err != nil {
		return err
	}

	filePathName = fmt.Sprint(dir, "/", l.object, ".", extension)

	if err = os.WriteFile(filePathName, data, 0666); err != nil {
		return err
	}

	return nil
}

func (l *LocalStorage) SetMineType(minetype string) *LocalStorage {
	l.mineType = minetype
	return l
}

func (l *LocalStorage) SetExtension(extension string) *LocalStorage {
	l.extension = extension
	return l
}

func (l *LocalStorage) Object(object string) *LocalStorage {
	l.object = object
	return l
}

func (l *LocalStorage) SetDirectory(dirs ...string) *LocalStorage {
	l.directories = dirs
	return l
}

func (l *LocalStorage) Bucket(bucket string) *LocalStorage {
	l.bucket = bucket
	return l
}

func (l *LocalStorage) Close() error {
	return nil
}

func (l *LocalStorage) mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (l *LocalStorage) getMineExtension(mine string) string {
	return l.extension
}

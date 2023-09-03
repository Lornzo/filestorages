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

func (l *LocalStorage) getDirectoryPath() string {
	return l.bucket + "/" + strings.Join(l.directories, "/")
}

func (l *LocalStorage) Upload(ctx context.Context, data []byte) error {

	var (
		dir          string = l.getDirectoryPath()
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

func (l *LocalStorage) Delete(ctx context.Context) error {

	var (
		dir          string = l.getDirectoryPath()
		extension    string = l.getMineExtension(l.mineType)
		err          error
		filePathName string
	)

	if extension == "" {
		return errors.New("not support mine type")
	}

	filePathName = fmt.Sprint(dir, "/", l.object, ".", extension)

	if err = os.Remove(filePathName); err != nil {
		return err
	}

	return nil
}

func (l *LocalStorage) SetMineType(minetype string) *LocalStorage {
	l.mineType = minetype
	l.SetExtension(l.getMineExtension(minetype))
	return l
}

func (l *LocalStorage) SetExtension(extension string) *LocalStorage {
	l.extension = extension
	return l
}

func (l *LocalStorage) GetExtension() string {
	return l.extension
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

package gcpstorage

import "cloud.google.com/go/storage"

type Writer struct {
	*storage.Writer
}

func (w *Writer) SetChunkSize(size int) {
	w.ChunkSize = size
}

func (w *Writer) Write(p []byte) (n int, err error) {
	return w.Writer.Write(p)
}

func (w *Writer) Close() error {
	return w.Writer.Close()
}

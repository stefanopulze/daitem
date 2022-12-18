package storage

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileStorage struct {
	dir string
}

func (f *FileStorage) Write(key string, value []byte) error {
	file, err := f.fileForWrite(key)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(value)
	return err
}

func (f *FileStorage) Read(key string) ([]byte, error) {
	file, err := f.fileForRead(key)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	return bytes, nil
}

func (f *FileStorage) Delete(key string) error {
	return os.Remove(f.filePathToFile(key))
}

func NewFileStorage(dir string) (Storage, error) {
	os.Mkdir(dir, 0777)

	return &FileStorage{
		dir: dir,
	}, nil
}

// Returns a string where invalid characters (e.g. colon ":" which is not allowed in file names on Window) are removed from fname
func removeInvalidFileNameCharacters(fname string) string {
	return strings.Replace(fname, ":", "", -1)
}

func (f *FileStorage) filePathToFile(file string) string {
	fname := removeInvalidFileNameCharacters(file)
	return filepath.Join(f.dir, fname)
}

func (f *FileStorage) fileForWrite(key string) (*os.File, error) {
	return os.OpenFile(f.filePathToFile(key), os.O_WRONLY|os.O_CREATE, 0666)
}

func (f *FileStorage) fileForRead(key string) (*os.File, error) {
	return os.OpenFile(f.filePathToFile(key), os.O_RDONLY, 0666)
}

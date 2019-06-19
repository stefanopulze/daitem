package storage

import (
	"strings"
)

type FileStorage struct {
}

func (FileStorage) Write(key string, value []byte) error {
	panic("implement me")
}

func (FileStorage) Read(key string) ([]byte, error) {
	panic("implement me")
}

func (FileStorage) Delete(key string) {
	panic("implement me")
}

func NewFileStorage(dir string) (Storage, error) {
	return FileStorage{}, nil
}

// Returns a string where invalid characters (e.g. colon ":" which is not allowed in file names on Window) are removed from fname
func removeInvalidFileNameCharacters(fname string) string {
	return strings.Replace(fname, ":", "", -1)
}

//func (f *FileStorage) fileForWrite(key string) (*os.File, error) {
//	return os.OpenFile(f.filePathToFile(key), os.O_WRONLY|os.O_CREATE, 0666)
//}
//
//func (f *FileStorage) fileForRead(key string) (*os.File, error) {
//	return os.OpenFile(f.filePathToFile(key), os.O_RDONLY, 0666)
//}

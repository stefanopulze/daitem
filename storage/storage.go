package storage

type Storage interface {
	Write(key string, value []byte) error

	Delete(key string)

	Read(key string) ([]byte, error)
}

package daitem

import (
	"github.com/stefanopulze/daitem/storage"
	"os"
)

type ClientOptions struct {
	Storage       *storage.Storage
	Username      string
	Password      string
	MasterCode    string
	CentralId     string
	TransmitterId string
}

func DefaultOptions(username string, password string, masterCode string) (*ClientOptions, error) {
	storage, err := storage.NewFileStorage(os.TempDir() + "/daitem")
	if err != nil {
		return nil, err
	}

	return &ClientOptions{
		Username:   username,
		Password:   password,
		MasterCode: masterCode,
		Storage:    &storage,
	}, nil
}

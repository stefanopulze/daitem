package storage

import (
	"os"
	"testing"
)

func TestNewFileStorage(t *testing.T) {
	s, e := NewFileStorage(os.TempDir() + "/daitem")
	if e != nil {
		t.Fatal(e)
	}

	if e := s.Write("Name", []byte("Stefano")); e != nil {
		t.Fatal(e)
	}

	if bytes, e := s.Read("Name"); e != nil {
		t.Fatal(e)
	} else {
		value := string(bytes)

		if value != "Stefano" {
			t.Fatal("Value error")
		}
	}

}

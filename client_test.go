package daitem

import (
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	options, _ := DefaultOptions("EMAIL", "PWD", "1234")

	client := NewClient(options)

	if status, err := client.Status(); err != nil {
		t.Fatal(err)
	} else {
		log.Printf("Current status: %v", status)
	}
}

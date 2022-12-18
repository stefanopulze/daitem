package daitem

import (
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	options, _ := DefaultOptions(
		"stefano.pulze87@gmail.com",
		"SP+daitem5715!*",
		"0715")

	client := NewClient(options)

	if status, err := client.Status(); err != nil {
		t.Fatal(err)
	} else {
		log.Printf("Current status: %v", status)
	}

	client.TurnAlarm(false)
}

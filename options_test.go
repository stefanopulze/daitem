package daitem

import "testing"

func TestDefaultOptions(t *testing.T) {
	options, err := DefaultOptions("mario", "password", "1234")

	if err != nil {
		t.Fatal(err)
	}

	if options.Username != "mario" {
		t.Fatal("Username not match")
	}

	if options.Password != "password" {
		t.Fatal("Password not match")
	}

	if options.MasterCode != "1234" {
		t.Fatal("MasterCode not match")
	}
}

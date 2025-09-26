package file

import (
	"fmt"
	"testing"
)

func TestNewHost_1(t *testing.T) {
	db := &DbUtils{
		JsonDb: &JsonDb{},
	}

	t.Run("test new host with empty location", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		host := &Host{
			Id:       1,
			Host:     "localhost",
			Location: "",
		}
		err := db.NewHost(host)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		if host.Location != "/" {
			t.Errorf("expected location to be '/', got %v", host.Location)
		}
	})

	t.Run("test new host with non-empty location", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		host := &Host{
			Id:       2,
			Host:     "localhost",
			Location: "/test",
		}
		err := db.NewHost(host)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		if host.Location != "/test" {
			t.Errorf("expected location to be '/test', got %v", host.Location)
		}
	})

	t.Run("test new host with existing id", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		host := &Host{
			Id:       1,
			Host:     "localhost",
			Location: "/test",
		}
		err := db.NewHost(host)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err.Error() != "host has exist" {
			t.Errorf("expected error message 'host has exist', got %v", err.Error())
		}
	})
}

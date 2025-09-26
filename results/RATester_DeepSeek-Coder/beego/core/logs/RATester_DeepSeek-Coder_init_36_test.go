package logs

import (
	"fmt"
	"testing"
)

func TestInit_36(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	Register(AdapterMail, newSMTPWriter)

	if _, ok := adapters[AdapterMail]; !ok {
		t.Errorf("Expected adapter %s to be registered", AdapterMail)
	}

	writer := adapters[AdapterMail]()
	if _, ok := writer.(Logger); !ok {
		t.Errorf("Expected writer to be of type Logger")
	}
}

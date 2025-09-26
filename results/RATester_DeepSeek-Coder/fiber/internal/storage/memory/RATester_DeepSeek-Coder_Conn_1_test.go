package memory

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	storage := &Storage{
		db: map[string]entry{
			"key1": {
				data:   []byte("value1"),
				expiry: 0,
			},
			"key2": {
				data:   []byte("value2"),
				expiry: 0,
			},
		},
		done:       make(chan struct{}),
		gcInterval: 0,
	}

	expected := storage.db
	result := storage.Conn()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

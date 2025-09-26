package memory

import (
	"fmt"
	"testing"
)

func TestDelete_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Storage{
		db: map[string]entry{
			"key1": {
				data: []byte("value1"),
			},
			"key2": {
				data: []byte("value2"),
			},
		},
	}
	key := "key1"
	err := s.Delete(key)
	if err != nil {
		t.Errorf("Delete() error = %v, wantErr nil", err)
	}
	if _, ok := s.db[key]; ok {
		t.Errorf("Delete() key = %v, want %v", ok, false)
	}
}

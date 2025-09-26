package orm

import (
	"fmt"
	"testing"
)

func TestBegin_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingOrm{}
	tx, err := d.Begin()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if tx != nil {
		t.Errorf("Expected nil, got %v", tx)
	}
}

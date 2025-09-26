package modules

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestTime_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &moduleAdapter{}
	if got := m.Time(); !reflect.DeepEqual(got, time.Time{}) {
		t.Errorf("Time() = %v, want %v", got, time.Time{})
	}
}

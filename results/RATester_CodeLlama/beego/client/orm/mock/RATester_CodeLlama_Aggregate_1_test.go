package mock

import (
	"fmt"
	"testing"
)

func TestAggregate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := &DoNothingQuerySetter{}
	s := "test"
	got := qs.Aggregate(s)
	if got != qs {
		t.Errorf("Aggregate() = %v, want %v", got, qs)
	}
}

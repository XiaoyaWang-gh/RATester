package mock

import (
	"fmt"
	"testing"
)

func TestFilterRaw_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	s := "s"
	s2 := "s2"
	d.FilterRaw(s, s2)
}

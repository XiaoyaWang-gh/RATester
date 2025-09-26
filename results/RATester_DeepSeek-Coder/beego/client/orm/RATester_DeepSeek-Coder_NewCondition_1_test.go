package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCondition_1(t *testing.T) {
	t.Run("should return a new condition", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := NewCondition()
		want := &Condition{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

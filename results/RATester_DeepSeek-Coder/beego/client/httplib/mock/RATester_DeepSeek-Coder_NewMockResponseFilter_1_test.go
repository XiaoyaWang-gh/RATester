package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewMockResponseFilter_1(t *testing.T) {
	t.Run("NewMockResponseFilter", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := NewMockResponseFilter()
		want := &MockResponseFilter{
			ms: make([]*Mock, 0, 1),
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

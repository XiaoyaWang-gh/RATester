package helpers

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestReaderToBytes_1(t *testing.T) {
	t.Parallel()

	t.Run("nil reader", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := ReaderToBytes(nil)
		want := []byte{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("empty reader", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		r := strings.NewReader("")
		got := ReaderToBytes(r)
		want := []byte{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("non-empty reader", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		r := strings.NewReader("hello, world")
		got := ReaderToBytes(r)
		want := []byte("hello, world")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

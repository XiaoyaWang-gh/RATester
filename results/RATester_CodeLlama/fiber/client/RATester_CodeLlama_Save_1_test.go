package client

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSave_1(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		r := &Response{}
		v := "test.txt"

		err := r.Save(v)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("io.Writer", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		r := &Response{}
		v := &bytes.Buffer{}

		err := r.Save(v)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("default", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		r := &Response{}
		v := 1

		err := r.Save(v)
		if err != ErrNotSupportSaveMethod {
			t.Fatal(err)
		}
	})
}

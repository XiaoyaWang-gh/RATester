package exif

import (
	"fmt"
	"testing"
)

func TestMarshalJSON_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := Tags{
		"foo": "bar",
	}
	b, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"foo":"bar"}` {
		t.Errorf("got %q, want %q", b, `{"foo":"bar"}`)
	}
}

package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalMsg_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &data{
		Data: map[string]any{
			"key1": "value1",
			"key2": 123,
		},
	}
	b, err := d.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}
	d2 := &data{}
	_, err = d2.UnmarshalMsg(b)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(d, d2) {
		t.Fatalf("expected %v, got %v", d, d2)
	}
}

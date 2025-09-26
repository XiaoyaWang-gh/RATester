package types

import (
	"fmt"
	"testing"
)

func TestMarshalJSON_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	q := &BandwidthQuantity{
		s: "100MB",
		i: 100 * 1024 * 1024,
	}
	b, err := q.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "\"100MB\"" {
		t.Fatalf("expect \"100MB\", but got %s", string(b))
	}
}

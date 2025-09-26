package echo

import (
	"fmt"
	"testing"
)

func TestMustBindWithDelimiter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Values []string
	}

	b := &ValueBinder{
		ValuesFunc: func(sourceParam string) []string {
			if sourceParam == "ids" {
				return []string{"1", "2"}
			}
			return nil
		},
	}

	ts := &testStruct{}
	b.MustBindWithDelimiter("ids", &ts.Values, ",")

	if len(ts.Values) != 2 || ts.Values[0] != "1" || ts.Values[1] != "2" {
		t.Errorf("Expected values to be [\"1\", \"2\"], got %v", ts.Values)
	}
}

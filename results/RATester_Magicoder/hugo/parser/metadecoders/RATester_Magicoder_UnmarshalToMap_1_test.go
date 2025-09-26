package metadecoders

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalToMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := Decoder{
		Delimiter:  ',',
		Comment:    '#',
		LazyQuotes: true,
	}

	data := []byte(`key1,value1
key2,value2
key3,value3`)

	expected := map[string]any{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	result, err := d.UnmarshalToMap(data, d.FormatFromContentString(string(data)))

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

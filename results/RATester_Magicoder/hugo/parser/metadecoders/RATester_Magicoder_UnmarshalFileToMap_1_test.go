package metadecoders

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestUnmarshalFileToMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	filename := "test.json"
	data := []byte(`{"key": "value"}`)

	err := afero.WriteFile(fs, filename, data, 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	d := Decoder{}
	result, err := d.UnmarshalFileToMap(fs, filename)
	if err != nil {
		t.Fatalf("Failed to unmarshal file to map: %v", err)
	}

	expected := map[string]any{"key": "value"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

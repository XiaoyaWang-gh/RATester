package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSuffixes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Type{
		SuffixesCSV: "jpg,jpeg",
	}

	if got := m.Suffixes(); !reflect.DeepEqual(got, []string{"jpg", "jpeg"}) {
		t.Errorf("Suffixes() = %v, want %v", got, []string{"jpg", "jpeg"})
	}
}

package attributes

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOptionsSlice_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &AttributesHolder{
		options: []Attribute{
			{Name: "foo", Value: "bar"},
			{Name: "baz", Value: "qux"},
		},
	}

	if diff := cmp.Diff(a.OptionsSlice(), []Attribute{
		{Name: "foo", Value: "bar"},
		{Name: "baz", Value: "qux"},
	}); diff != "" {
		t.Fatalf("unexpected result (-want +got):\n%s", diff)
	}
}

package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepCopy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	in := &Domain{
		Main: "foo.com",
		SANs: []string{"bar.com", "baz.com"},
	}
	out := in.DeepCopy()
	if out == nil {
		t.Fatal("expected out to not be nil")
	}
	if out.Main != in.Main {
		t.Errorf("expected out.Main to be %s but got %s", in.Main, out.Main)
	}
	if !reflect.DeepEqual(out.SANs, in.SANs) {
		t.Errorf("expected out.SANs to be %v but got %v", in.SANs, out.SANs)
	}
}

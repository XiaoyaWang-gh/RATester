package security

import (
	"fmt"
	"testing"
)

func TestMarshalJSON_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := Whitelist{
		acceptNone: true,
	}

	b, err := w.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != `"none"` {
		t.Errorf("expected %q, got %q", `"none"`, string(b))
	}
}

package internal

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMarshalJSON_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	namespaces := TemplateFuncsNamespaces{
		{
			Name: "strings",
		},
		{
			Name: "lang",
		},
	}

	b, err := namespaces.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b, []byte(`{"strings":null,"lang":null}`)) {
		t.Fatalf("unexpected result: %s", b)
	}
}

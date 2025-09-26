package binding

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestBind_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := msgpackBinding{}

	type testStruct struct {
		Field1 string `msgpack:"field1"`
		Field2 int    `msgpack:"field2"`
	}

	testObj := testStruct{}

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte{0x82, 0xa6, 'f', 'i', 'e', 'l', 'd', '1', 0xa6, 'f', 'i', 'e', 'l', 'd', '2', 0x01}))
	if err != nil {
		t.Fatal(err)
	}

	err = b.Bind(req, &testObj)
	if err != nil {
		t.Fatal(err)
	}

	if testObj.Field1 != "field1" {
		t.Errorf("Expected Field1 to be 'field1', got %s", testObj.Field1)
	}

	if testObj.Field2 != 1 {
		t.Errorf("Expected Field2 to be 1, got %d", testObj.Field2)
	}
}

package forwardedheaders

import (
	"fmt"
	"testing"
)

func TestDel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	header := make(unsafeHeader)
	header.Set("key1", "value1")
	header.Set("key2", "value2")

	header.Del("key1")

	if _, ok := header["key1"]; ok {
		t.Error("Expected key1 to be deleted")
	}

	if _, ok := header["key2"]; !ok {
		t.Error("Expected key2 to still exist")
	}
}

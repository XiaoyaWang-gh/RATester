package converter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := Bytes{0x01, 0x02, 0x03, 0x04, 0x05}
	if !reflect.DeepEqual(b.Bytes(), []byte{0x01, 0x02, 0x03, 0x04, 0x05}) {
		t.Errorf("Bytes() = %v, want %v", b.Bytes(), []byte{0x01, 0x02, 0x03, 0x04, 0x05})
	}
}

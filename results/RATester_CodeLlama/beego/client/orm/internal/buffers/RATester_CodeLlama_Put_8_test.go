package buffers

import (
	"fmt"
	"testing"

	"github.com/valyala/bytebufferpool"
)

func TestPut_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	bf := bytebufferpool.Get()
	defer Put(bf)

	// Act
	bf.WriteString("Hello, World!")

	// Assert
	if bf.String() != "Hello, World!" {
		t.Errorf("bf.String() = %v, want %v", bf.String(), "Hello, World!")
	}
}

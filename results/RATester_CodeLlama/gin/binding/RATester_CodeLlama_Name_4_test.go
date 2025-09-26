package binding

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestName_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	binding := protobufBinding{}
	assert.Equal(t, "protobuf", binding.Name())
}

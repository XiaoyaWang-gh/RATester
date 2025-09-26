package csrf

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMarshalMsg_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	z := storageManager{}
	b := []byte{}

	// When
	o, err := z.MarshalMsg(b)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, []byte{0x80}, o)
}

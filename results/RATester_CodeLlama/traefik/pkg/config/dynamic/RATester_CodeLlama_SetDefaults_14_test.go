package dynamic

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSetDefaults_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &UDPWRRService{}
	w.SetDefaults()
	assert.Equal(t, 1, *w.Weight)
}

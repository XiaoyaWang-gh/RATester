package ingress

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSetDefaults_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &ServiceIng{}
	s.SetDefaults()
	assert.Equal(t, true, *s.PassHostHeader)
}

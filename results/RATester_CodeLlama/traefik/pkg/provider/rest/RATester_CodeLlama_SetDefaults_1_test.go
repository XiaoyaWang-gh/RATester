package rest

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSetDefaults_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	p.SetDefaults()
	assert.Equal(t, false, p.Insecure)
}

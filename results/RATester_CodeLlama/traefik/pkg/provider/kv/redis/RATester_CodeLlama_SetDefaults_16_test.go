package redis

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSetDefaults_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	p.SetDefaults()
	assert.Equal(t, []string{"127.0.0.1:6379"}, p.Endpoints)
}

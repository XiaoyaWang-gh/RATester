package consul

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSetDefaults_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ProviderBuilder{}
	p.SetDefaults()
	assert.Equal(t, []string{"127.0.0.1:8500"}, p.Endpoints)
}

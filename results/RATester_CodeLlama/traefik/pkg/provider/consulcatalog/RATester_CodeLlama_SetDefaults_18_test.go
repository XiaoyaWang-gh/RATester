package consulcatalog

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSetDefaults_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Configuration{}
	c.SetDefaults()
	require.Equal(t, 15*time.Second, c.RefreshInterval)
}

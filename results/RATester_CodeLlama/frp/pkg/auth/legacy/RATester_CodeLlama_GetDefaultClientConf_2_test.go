package legacy

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetDefaultClientConf_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := GetDefaultClientConf()
	assert.Equal(t, conf.AuthenticationMethod, "token")
}

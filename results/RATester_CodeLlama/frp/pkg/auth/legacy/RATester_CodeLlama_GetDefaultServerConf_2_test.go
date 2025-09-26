package legacy

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetDefaultServerConf_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := GetDefaultServerConf()
	assert.Equal(t, conf.AuthenticationMethod, "token")
	assert.Equal(t, conf.AuthenticateHeartBeats, false)
	assert.Equal(t, conf.AuthenticateNewWorkConns, false)
}

package legacy

import (
	"fmt"
	"testing"

	legacyauth "github.com/fatedier/frp/pkg/auth/legacy"
	"gotest.tools/assert"
)

func TestGetDefaultClientConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := GetDefaultClientConf()
	assert.Equal(t, conf.ClientConfig, legacyauth.GetDefaultClientConf())
	assert.Equal(t, conf.TCPMux, true)
	assert.Equal(t, conf.LoginFailExit, true)
	assert.Equal(t, conf.Protocol, "tcp")
	assert.Equal(t, conf.Start, []string{})
	assert.Equal(t, conf.TLSEnable, true)
	assert.Equal(t, conf.DisableCustomTLSFirstByte, true)
	assert.Equal(t, conf.Metas, map[string]string{})
	assert.Equal(t, conf.IncludeConfigFiles, []string{})
}

package config

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/spf13/cobra"
)

func TestRegisterProxyFlags_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := &cobra.Command{}
	c := &v1.TCPProxyConfig{}
	RegisterProxyFlags(cmd, c)

	if cmd.Flags().Lookup("remote_port") == nil {
		t.Error("remote_port flag not found")
	}
}

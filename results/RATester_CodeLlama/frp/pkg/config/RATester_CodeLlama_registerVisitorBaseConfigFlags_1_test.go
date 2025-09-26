package config

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/spf13/cobra"
)

func TestRegisterVisitorBaseConfigFlags_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := &cobra.Command{}
	c := &v1.VisitorBaseConfig{}
	registerVisitorBaseConfigFlags(cmd, c)
	if cmd.Flags().Lookup("visitor_name") == nil {
		t.Error("visitor_name flag not found")
	}
	if cmd.Flags().Lookup("ue") == nil {
		t.Error("ue flag not found")
	}
	if cmd.Flags().Lookup("uc") == nil {
		t.Error("uc flag not found")
	}
	if cmd.Flags().Lookup("sk") == nil {
		t.Error("sk flag not found")
	}
	if cmd.Flags().Lookup("server_name") == nil {
		t.Error("server_name flag not found")
	}
	if cmd.Flags().Lookup("bind_addr") == nil {
		t.Error("bind_addr flag not found")
	}
	if cmd.Flags().Lookup("bind_port") == nil {
		t.Error("bind_port flag not found")
	}
}

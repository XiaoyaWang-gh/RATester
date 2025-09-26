package config

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/spf13/cobra"
)

func TestRegisterServerConfigFlags_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := &cobra.Command{}
	c := &v1.ServerConfig{}
	RegisterServerConfigFlags(cmd, c)
	// TODO: check flags
}

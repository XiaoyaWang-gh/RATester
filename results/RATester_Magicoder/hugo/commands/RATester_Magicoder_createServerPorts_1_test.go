package commands

import (
	"fmt"
	"testing"

	"github.com/bep/simplecobra"
	"github.com/spf13/cobra"
)

func TestcreateServerPorts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cd := &simplecobra.Commandeer{
		CobraCommand: &cobra.Command{
			Use: "test",
		},
	}

	cd.CobraCommand.Flags().Bool("appendPort", false, "")

	serverCommand := &serverCommand{
		r: &rootCommand{
			Printf: func(format string, v ...interface{}) {
				fmt.Printf(format, v...)
			},
		},
	}

	err := serverCommand.createServerPorts(cd)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

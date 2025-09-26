package sub

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func TestInit_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rootCmd := &cobra.Command{}
	verifyCmd := &cobra.Command{}
	rootCmd.AddCommand(verifyCmd)
	if rootCmd.Commands()[0] != verifyCmd {
		t.Errorf("rootCmd.Commands()[0] != verifyCmd")
	}
}

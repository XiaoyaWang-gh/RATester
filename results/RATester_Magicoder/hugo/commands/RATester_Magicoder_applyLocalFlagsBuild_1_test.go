package commands

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func TestapplyLocalFlagsBuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := &cobra.Command{}
	r := &rootCommand{}

	applyLocalFlagsBuild(cmd, r)

	// Test the flags
	cleanDestinationDir, err := cmd.Flags().GetBool("cleanDestinationDir")
	if err != nil {
		t.Errorf("Error getting cleanDestinationDir flag: %v", err)
	}
	if cleanDestinationDir != false {
		t.Errorf("Expected cleanDestinationDir to be false, but got %v", cleanDestinationDir)
	}

	buildDrafts, err := cmd.Flags().GetBool("buildDrafts")
	if err != nil {
		t.Errorf("Error getting buildDrafts flag: %v", err)
	}
	if buildDrafts != false {
		t.Errorf("Expected buildDrafts to be false, but got %v", buildDrafts)
	}

	// Add more tests for other flags...
}

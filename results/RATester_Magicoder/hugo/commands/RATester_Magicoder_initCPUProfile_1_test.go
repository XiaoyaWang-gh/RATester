package commands

import (
	"fmt"
	"os"
	"testing"
)

func TestinitCPUProfile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hb := &hugoBuilder{
		r: &rootCommand{
			cpuprofile: "test.prof",
		},
	}

	cleanup, err := hb.initCPUProfile()
	if err != nil {
		t.Fatalf("initCPUProfile failed: %v", err)
	}

	// Perform some operations that should be profiled

	cleanup()

	// Check if the CPU profile file exists
	_, err = os.Stat("test.prof")
	if os.IsNotExist(err) {
		t.Fatal("CPU profile file not found")
	}

	// Clean up the profile file
	err = os.Remove("test.prof")
	if err != nil {
		t.Fatalf("Failed to remove CPU profile file: %v", err)
	}
}

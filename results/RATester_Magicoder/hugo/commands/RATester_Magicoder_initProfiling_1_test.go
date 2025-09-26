package commands

import (
	"fmt"
	"os"
	"testing"
)

func TestinitProfiling_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hb := &hugoBuilder{
		r: &rootCommand{
			cpuprofile: "cpu.prof",
		},
	}

	stopProfiling, err := hb.initProfiling()
	if err != nil {
		t.Errorf("initProfiling() returned an error: %v", err)
	}

	if stopProfiling == nil {
		t.Error("initProfiling() did not return a stop function")
	}

	stopProfiling()

	if _, err := os.Stat("cpu.prof"); os.IsNotExist(err) {
		t.Error("CPU profile was not created")
	}

	os.Remove("cpu.prof")
}

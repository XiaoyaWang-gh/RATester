package commands

import (
	"fmt"
	"os"
	"testing"
)

func TestinitMutexProfile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hugoBuilder := &hugoBuilder{
		r: &rootCommand{
			mutexprofile: "test.prof",
		},
	}

	cleanup, err := hugoBuilder.initMutexProfile()
	if err != nil {
		t.Fatalf("initMutexProfile() failed: %v", err)
	}

	if cleanup == nil {
		t.Fatal("initMutexProfile() returned nil cleanup function")
	}

	cleanup()

	if _, err := os.Stat("test.prof"); os.IsNotExist(err) {
		t.Fatal("initMutexProfile() did not create mutex profile file")
	}

	os.Remove("test.prof")
}

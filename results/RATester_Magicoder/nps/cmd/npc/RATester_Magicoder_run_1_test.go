package main

import (
	"fmt"
	"os"
	"testing"
)

func Testrun_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Setup
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "-serverAddr=localhost", "-verifyKey=testKey", "-connType=tcp", "-password=testPassword", "-localType=testType", "-target=testTarget", "-localPort=8080"}

	// Call the function under test
	run()

	// Assertions
	// You can add assertions here to check if the function behaves as expected
	// For example:
	// assert.Equal(t, expected, actual, "The function run() should return the expected value")
}

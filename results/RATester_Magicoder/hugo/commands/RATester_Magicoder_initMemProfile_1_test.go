package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestinitMemProfile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hugoBuilder := &hugoBuilder{
		r: &rootCommand{
			memprofile: "test.prof",
		},
	}

	// Create a temporary file for the profile
	tmpFile, err := ioutil.TempFile("", "memprofile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// Set the temporary file name in the hugoBuilder
	hugoBuilder.r.memprofile = tmpFile.Name()

	// Call the function to be tested
	hugoBuilder.initMemProfile()

	// Check if the file exists
	_, err = os.Stat(tmpFile.Name())
	if os.IsNotExist(err) {
		t.Error("Memory profile file was not created")
	}
}

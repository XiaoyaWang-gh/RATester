package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGenKVDynConfDoc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	outputFile := "test.txt"
	genKVDynConfDoc(outputFile)

	_, err := os.Stat(outputFile)
	if os.IsNotExist(err) {
		t.Errorf("File %s does not exist", outputFile)
	}

	os.Remove(outputFile)
}

package main

import (
	"fmt"
	"os"
	"testing"
)

func Testmain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"test"}

	main()
}

package commands

import (
	"fmt"
	"os"
	"testing"
)

func Testbuild_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hugoBuilder := &hugoBuilder{
		r: &rootCommand{
			Printf: func(format string, v ...interface{}) {
				fmt.Printf(format, v...)
			},
			Println: func(a ...interface{}) {
				fmt.Println(a...)
			},
			Out: os.Stdout,
		},
	}

	err := hugoBuilder.build()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

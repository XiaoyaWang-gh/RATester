package admin

import (
	"fmt"
	"io"
	"testing"
)

func TestGetCPUProfile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var w io.Writer
	GetCPUProfile(w)
}

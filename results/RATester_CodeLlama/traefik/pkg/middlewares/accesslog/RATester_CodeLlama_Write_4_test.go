package accesslog

import (
	"fmt"
	"os"
	"testing"
)

func TestWrite_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := noopCloser{File: os.Stdout}
	p := []byte("Hello, world!")
	n.Write(p)
}

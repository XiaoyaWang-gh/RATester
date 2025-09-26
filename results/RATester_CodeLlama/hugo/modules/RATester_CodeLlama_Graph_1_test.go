package modules

import (
	"fmt"
	"os"
	"testing"
)

func TestGraph_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	err := c.Graph(os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}

package group

import (
	"fmt"
	"testing"
)

func TestChooseEndpoint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &HTTPGroup{
		group: "test",
	}
	name, err := g.chooseEndpoint()
	if err != nil {
		t.Error(err)
	}
	if name != "" {
		t.Error("name should be empty")
	}
}

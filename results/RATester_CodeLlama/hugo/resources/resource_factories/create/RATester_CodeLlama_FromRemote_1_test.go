package create

import (
	"fmt"
	"testing"
)

func TestFromRemote_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	uri := "https://example.com/foo.txt"
	optionsm := map[string]any{}

	_, err := c.FromRemote(uri, optionsm)
	if err != nil {
		t.Fatal(err)
	}
}

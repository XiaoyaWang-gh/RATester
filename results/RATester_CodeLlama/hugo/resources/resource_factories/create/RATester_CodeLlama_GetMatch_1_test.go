package create

import (
	"fmt"
	"testing"
)

func TestGetMatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	res, err := c.GetMatch("")
	if err != nil {
		t.Errorf("GetMatch() error = %v", err)
		return
	}
	if res == nil {
		t.Errorf("GetMatch() res = nil")
		return
	}
}

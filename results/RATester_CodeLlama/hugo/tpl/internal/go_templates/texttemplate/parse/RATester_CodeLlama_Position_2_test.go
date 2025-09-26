package parse

import (
	"fmt"
	"testing"
)

func TestPosition_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Pos(1)
	if got := p.Position(); got != p {
		t.Errorf("Position() = %v, want %v", got, p)
	}
}

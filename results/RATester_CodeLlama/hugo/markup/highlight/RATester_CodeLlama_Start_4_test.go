package highlight

import (
	"fmt"
	"testing"
)

func TestStart_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := startEnd{}
	code := true
	styleAttr := "styleAttr"
	want := "start"
	got := s.Start(code, styleAttr)
	if got != want {
		t.Errorf("Start() = %v, want %v", got, want)
	}
}

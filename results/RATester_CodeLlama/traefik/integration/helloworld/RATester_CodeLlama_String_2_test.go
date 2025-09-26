package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestString_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &StreamExampleReply{
		Data: "data",
	}
	want := "data"
	if got := m.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

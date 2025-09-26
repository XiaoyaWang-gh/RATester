package alils

import (
	"fmt"
	"testing"
)

func TestSovLog_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input uint64
		want  int
	}

	tests := []test{
		{input: 0, want: 1},
		{input: 1, want: 1},
		{input: 128, want: 2},
		{input: 129, want: 3},
		{input: 16384, want: 5},
		{input: 16385, want: 6},
		{input: 18446744073709551615, want: 64},
	}

	for _, tt := range tests {
		got := sovLog(tt.input)
		if got != tt.want {
			t.Errorf("sovLog(%d) = %d; want %d", tt.input, got, tt.want)
		}
	}
}

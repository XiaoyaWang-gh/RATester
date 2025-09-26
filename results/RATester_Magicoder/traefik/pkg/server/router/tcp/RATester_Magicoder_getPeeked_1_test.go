package tcp

import (
	"bufio"
	"fmt"
	"testing"
)

func TestGetPeeked_1(t *testing.T) {
	type args struct {
		br *bufio.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getPeeked(tt.args.br); got != tt.want {
				t.Errorf("getPeeked() = %v, want %v", got, tt.want)
			}
		})
	}
}

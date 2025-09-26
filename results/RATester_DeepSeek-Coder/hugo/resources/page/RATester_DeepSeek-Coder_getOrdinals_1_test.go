package page

import (
	"fmt"
	"testing"
)

func TestGetOrdinals_1(t *testing.T) {
	type args struct {
		p1 Page
		p2 Page
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
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

			got, got1 := getOrdinals(tt.args.p1, tt.args.p2)
			if got != tt.want {
				t.Errorf("getOrdinals() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getOrdinals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

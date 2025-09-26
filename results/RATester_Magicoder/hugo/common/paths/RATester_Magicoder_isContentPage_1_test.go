package paths

import (
	"fmt"
	"testing"
)

func TestisContentPage_1(t *testing.T) {
	type args struct {
		p *Path
	}
	tests := []struct {
		name string
		args args
		want bool
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

			if got := tt.args.p.isContentPage(); got != tt.want {
				t.Errorf("isContentPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

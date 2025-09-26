package page

import (
	"fmt"
	"testing"
)

func TestExpandPattern_1(t *testing.T) {
	type args struct {
		pattern string
		p       Page
	}
	tests := []struct {
		name    string
		l       PermalinkExpander
		args    args
		want    string
		wantErr bool
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

			got, err := tt.l.ExpandPattern(tt.args.pattern, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpandPattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExpandPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

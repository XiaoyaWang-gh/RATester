package postpub

import (
	"fmt"
	"testing"
)

func TestGetFieldString_1(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name  string
		r     *PostPublishResource
		args  args
		want  string
		want1 bool
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

			got, got1 := tt.r.GetFieldString(tt.args.pattern)
			if got != tt.want {
				t.Errorf("GetFieldString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetFieldString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

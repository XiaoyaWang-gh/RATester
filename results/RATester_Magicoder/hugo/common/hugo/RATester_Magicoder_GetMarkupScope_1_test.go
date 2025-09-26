package hugo

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMarkupScope_1(t *testing.T) {
	type args struct {
		ctx context.Context
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

			if got := GetMarkupScope(tt.args.ctx); got != tt.want {
				t.Errorf("GetMarkupScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

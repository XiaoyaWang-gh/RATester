package binding

import (
	"fmt"
	"testing"
)

func TestmapHeader_1(t *testing.T) {
	type args struct {
		ptr any
		h   map[string][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				ptr: &struct{}{},
				h:   map[string][]string{"key": {"value"}},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				ptr: &struct{}{},
				h:   map[string][]string{"key": {"value"}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := mapHeader(tt.args.ptr, tt.args.h); (err != nil) != tt.wantErr {
				t.Errorf("mapHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

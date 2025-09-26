package binding

import (
	"fmt"
	"testing"
)

func TestmapURI_1(t *testing.T) {
	type args struct {
		ptr any
		m   map[string][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				ptr: &struct{ FieldName string }{},
				m:   map[string][]string{"FieldName": {"value"}},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				ptr: &struct{ FieldName string }{},
				m:   map[string][]string{"FieldName": {"value"}},
			},
			wantErr: true,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := mapURI(tt.args.ptr, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("mapURI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

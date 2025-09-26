package binder

import (
	"fmt"
	"testing"
)

func TestparseToStruct_1(t *testing.T) {
	type args struct {
		aliasTag string
		out      any
		data     map[string][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				aliasTag: "aliasTag",
				out:      &struct{}{},
				data:     map[string][]string{"field name": {"value"}},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				aliasTag: "aliasTag",
				out:      &struct{}{},
				data:     map[string][]string{"field name": {"value"}},
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

			if err := parseToStruct(tt.args.aliasTag, tt.args.out, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("parseToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

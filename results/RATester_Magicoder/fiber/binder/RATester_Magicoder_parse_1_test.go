package binder

import (
	"fmt"
	"testing"
)

func Testparse_1(t *testing.T) {
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
			name: "Test case 1",
			args: args{
				aliasTag: "test_tag",
				out:      map[string]string{},
				data:     map[string][]string{"field name": {"value"}},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				aliasTag: "test_tag",
				out:      struct{}{},
				data:     map[string][]string{"field name": {"value"}},
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

			if err := parse(tt.args.aliasTag, tt.args.out, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

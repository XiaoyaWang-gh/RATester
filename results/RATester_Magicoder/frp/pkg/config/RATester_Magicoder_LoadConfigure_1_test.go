package config

import (
	"fmt"
	"testing"
)

func TestLoadConfigure_1(t *testing.T) {
	type args struct {
		b      []byte
		c      any
		strict bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				b:      []byte("field_name: value"),
				c:      &struct{ FieldName string }{},
				strict: false,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				b:      []byte("field_name: value"),
				c:      &struct{ FieldName string }{},
				strict: true,
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

			if err := LoadConfigure(tt.args.b, tt.args.c, tt.args.strict); (err != nil) != tt.wantErr {
				t.Errorf("LoadConfigure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

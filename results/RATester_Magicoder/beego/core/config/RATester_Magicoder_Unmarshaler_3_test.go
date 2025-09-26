package config

import (
	"fmt"
	"testing"
)

func TestUnmarshaler_3(t *testing.T) {
	type args struct {
		prefix string
		obj    interface{}
		opt    []DecodeOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				prefix: "test_prefix",
				obj:    &struct{}{},
				opt:    []DecodeOption{},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				prefix: "test_prefix",
				obj:    &struct{}{},
				opt:    []DecodeOption{},
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

			if err := Unmarshaler(tt.args.prefix, tt.args.obj, tt.args.opt...); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshaler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

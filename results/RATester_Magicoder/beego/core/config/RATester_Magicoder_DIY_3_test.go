package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDIY_3(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				key: "test_key",
			},
			want:    "test_value",
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				key: "invalid_key",
			},
			want:    nil,
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

			got, err := DIY(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DIY() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DIY() = %v, want %v", got, tt.want)
			}
		})
	}
}

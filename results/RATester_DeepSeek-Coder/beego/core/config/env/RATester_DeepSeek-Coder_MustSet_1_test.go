package env

import (
	"fmt"
	"testing"
)

func TestMustSet_1(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				key:   "TEST_KEY",
				value: "TEST_VALUE",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				key:   "",
				value: "TEST_VALUE",
			},
			wantErr: true,
		},
		{
			name: "Test case 3",
			args: args{
				key:   "TEST_KEY",
				value: "",
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				key:   "",
				value: "",
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

			if err := MustSet(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("MustSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

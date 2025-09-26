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
				key:   "key1",
				value: "value1",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				key:   "key2",
				value: "value2",
			},
			wantErr: false,
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

			if err := MustSet(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("MustSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

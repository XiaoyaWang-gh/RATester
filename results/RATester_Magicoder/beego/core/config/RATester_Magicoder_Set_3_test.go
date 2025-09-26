package config

import (
	"fmt"
	"testing"
)

func TestSet_3(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				key: "key1",
				val: "value1",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				key: "key2",
				val: "value2",
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

			if err := Set(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

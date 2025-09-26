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
			name: "TestSet_0",
			args: args{
				key: "key0",
				val: "val0",
			},
			wantErr: false,
		},
		{
			name: "TestSet_1",
			args: args{
				key: "key1",
				val: "val1",
			},
			wantErr: false,
		},
		{
			name: "TestSet_2",
			args: args{
				key: "key2",
				val: "val2",
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

			if err := Set(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package config

import (
	"fmt"
	"testing"
)

func TestSet_4(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name    string
		c       *fakeConfigContainer
		args    args
		wantErr bool
	}{
		{
			name: "TestSet_0",
			c: &fakeConfigContainer{
				data: make(map[string]string),
			},
			args: args{
				key: "key",
				val: "value",
			},
			wantErr: false,
		},
		{
			name: "TestSet_1",
			c: &fakeConfigContainer{
				data: make(map[string]string),
			},
			args: args{
				key: "",
				val: "value",
			},
			wantErr: true,
		},
		{
			name: "TestSet_2",
			c: &fakeConfigContainer{
				data: make(map[string]string),
			},
			args: args{
				key: "key",
				val: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := tt.c
			if err := c.Set(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("fakeConfigContainer.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

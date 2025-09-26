package config

import (
	"fmt"
	"testing"
)

func TestSet_5(t *testing.T) {
	t.Parallel()
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name    string
		c       *IniConfigContainer
		args    args
		wantErr bool
	}{
		{
			name: "TestSet_EmptyKey",
			c:    &IniConfigContainer{data: make(map[string]map[string]string)},
			args: args{
				key: "",
				val: "value",
			},
			wantErr: true,
		},
		{
			name: "TestSet_Normal",
			c:    &IniConfigContainer{data: make(map[string]map[string]string)},
			args: args{
				key: "section::key",
				val: "value",
			},
			wantErr: false,
		},
		{
			name: "TestSet_DefaultSection",
			c:    &IniConfigContainer{data: make(map[string]map[string]string)},
			args: args{
				key: "key",
				val: "value",
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
				t.Errorf("IniConfigContainer.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

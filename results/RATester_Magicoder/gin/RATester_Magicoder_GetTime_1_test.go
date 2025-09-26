package gin

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetTime_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		c    *Context
		args args
		want time.Time
	}{
		{
			name: "TestGetTime_0",
			c: &Context{
				Keys: map[string]any{
					"key": time.Now(),
				},
			},
			args: args{
				key: "key",
			},
			want: time.Now(),
		},
		{
			name: "TestGetTime_1",
			c: &Context{
				Keys: map[string]any{
					"key": "not a time",
				},
			},
			args: args{
				key: "key",
			},
			want: time.Time{},
		},
		{
			name: "TestGetTime_2",
			c: &Context{
				Keys: map[string]any{
					"key": nil,
				},
			},
			args: args{
				key: "key",
			},
			want: time.Time{},
		},
		{
			name: "TestGetTime_3",
			c: &Context{
				Keys: map[string]any{
					"key": time.Now(),
				},
			},
			args: args{
				key: "not_exist",
			},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.GetTime(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

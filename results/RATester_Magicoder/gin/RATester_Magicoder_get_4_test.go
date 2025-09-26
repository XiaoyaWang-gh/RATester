package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func Testget_4(t *testing.T) {
	type args struct {
		m   map[string][]string
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]string
		want1 bool
	}{
		{
			name: "Test Case 1",
			args: args{
				m: map[string][]string{
					"key[subkey]": {"value"},
				},
				key: "key",
			},
			want:  map[string]string{"subkey": "value"},
			want1: true,
		},
		{
			name: "Test Case 2",
			args: args{
				m: map[string][]string{
					"key[subkey]": {"value"},
				},
				key: "otherkey",
			},
			want:  map[string]string{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{}
			got, got1 := c.get(tt.args.m, tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

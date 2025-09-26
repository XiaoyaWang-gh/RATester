package env

import (
	"fmt"
	"testing"
)

func TestGet_22(t *testing.T) {
	type args struct {
		key    string
		defVal string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				key:    "key1",
				defVal: "default1",
			},
			want: "value1",
		},
		{
			name: "Test case 2",
			args: args{
				key:    "key2",
				defVal: "default2",
			},
			want: "default2",
		},
		{
			name: "Test case 3",
			args: args{
				key:    "key3",
				defVal: "default3",
			},
			want: "default3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Get(tt.args.key, tt.args.defVal); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

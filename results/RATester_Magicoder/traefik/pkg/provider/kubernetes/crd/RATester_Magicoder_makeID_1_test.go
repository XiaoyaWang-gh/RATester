package crd

import (
	"fmt"
	"testing"
)

func TestMakeID_1(t *testing.T) {
	type args struct {
		namespace string
		name      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				namespace: "test",
				name:      "name",
			},
			want: "test-name",
		},
		{
			name: "Test case 2",
			args: args{
				namespace: "",
				name:      "name",
			},
			want: "name",
		},
		{
			name: "Test case 3",
			args: args{
				namespace: "test",
				name:      "",
			},
			want: "test",
		},
		{
			name: "Test case 4",
			args: args{
				namespace: "",
				name:      "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := makeID(tt.args.namespace, tt.args.name); got != tt.want {
				t.Errorf("makeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

package crd

import (
	"fmt"
	"testing"
)

func TestIsNamespaceAllowed_1(t *testing.T) {
	type args struct {
		allowCrossNamespace bool
		parentNamespace     string
		namespace           string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				allowCrossNamespace: true,
				parentNamespace:     "parent",
				namespace:           "child",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				allowCrossNamespace: false,
				parentNamespace:     "parent",
				namespace:           "child",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				allowCrossNamespace: true,
				parentNamespace:     "parent",
				namespace:           "parent",
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				allowCrossNamespace: false,
				parentNamespace:     "parent",
				namespace:           "parent",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isNamespaceAllowed(tt.args.allowCrossNamespace, tt.args.parentNamespace, tt.args.namespace); got != tt.want {
				t.Errorf("isNamespaceAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

package runtime

import (
	"fmt"
	"testing"
)

func TestGetQualifiedName_2(t *testing.T) {
	type args struct {
		provider    string
		elementName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with valid provider and elementName",
			args: args{
				provider:    "testProvider",
				elementName: "testElement",
			},
			want: "testElement@testProvider",
		},
		{
			name: "Test with valid provider and elementName with @",
			args: args{
				provider:    "testProvider",
				elementName: "testElement@testProvider",
			},
			want: "testElement@testProvider",
		},
		{
			name: "Test with empty provider",
			args: args{
				provider:    "",
				elementName: "testElement",
			},
			want: "testElement@",
		},
		{
			name: "Test with empty elementName",
			args: args{
				provider:    "testProvider",
				elementName: "",
			},
			want: "@testProvider",
		},
		{
			name: "Test with empty provider and elementName",
			args: args{
				provider:    "",
				elementName: "",
			},
			want: "@",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getQualifiedName(tt.args.provider, tt.args.elementName); got != tt.want {
				t.Errorf("getQualifiedName() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			name: "Test case 1",
			args: args{
				provider:    "provider1",
				elementName: "element1",
			},
			want: "element1@provider1",
		},
		{
			name: "Test case 2",
			args: args{
				provider:    "provider2",
				elementName: "element2@provider2",
			},
			want: "element2@provider2",
		},
		{
			name: "Test case 3",
			args: args{
				provider:    "provider3",
				elementName: "element3@provider3@provider3",
			},
			want: "element3@provider3@provider3",
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

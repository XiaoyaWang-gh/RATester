package provider

import (
	"fmt"
	"testing"
)

func TestMakeQualifiedName_1(t *testing.T) {
	type args struct {
		providerName string
		elementName  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				providerName: "provider1",
				elementName:  "element1",
			},
			want: "element1@provider1",
		},
		{
			name: "Test case 2",
			args: args{
				providerName: "provider2",
				elementName:  "element2",
			},
			want: "element2@provider2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MakeQualifiedName(tt.args.providerName, tt.args.elementName); got != tt.want {
				t.Errorf("MakeQualifiedName() = %v, want %v", got, tt.want)
			}
		})
	}
}

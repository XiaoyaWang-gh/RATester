package runtime

import (
	"fmt"
	"testing"
)

func TestGetProviderName_1(t *testing.T) {
	type args struct {
		elementName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{elementName: "element@provider"},
			want: "provider",
		},
		{
			name: "Test case 2",
			args: args{elementName: "element"},
			want: "",
		},
		{
			name: "Test case 3",
			args: args{elementName: "@provider"},
			want: "provider",
		},
		{
			name: "Test case 4",
			args: args{elementName: "element@provider@extra"},
			want: "provider@extra",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getProviderName(tt.args.elementName); got != tt.want {
				t.Errorf("getProviderName() = %v, want %v", got, tt.want)
			}
		})
	}
}

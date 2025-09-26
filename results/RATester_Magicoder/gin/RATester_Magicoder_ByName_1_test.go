package gin

import (
	"fmt"
	"testing"
)

func TestByName_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		ps   Params
		args args
		want string
	}{
		{
			name: "Test case 1",
			ps:   Params{{"name", "value"}},
			args: args{"name"},
			want: "value",
		},
		{
			name: "Test case 2",
			ps:   Params{{"name", "value"}},
			args: args{"invalid"},
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

			if got := tt.ps.ByName(tt.args.name); got != tt.want {
				t.Errorf("ByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

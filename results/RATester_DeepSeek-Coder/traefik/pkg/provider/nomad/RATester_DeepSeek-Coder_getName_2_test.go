package nomad

import (
	"fmt"
	"testing"
)

func TestGetName_2(t *testing.T) {
	type args struct {
		i item
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with canary",
			args: args{
				i: item{
					Name:      "test",
					Tags:      []string{"tag1", "tag2"},
					ExtraConf: configuration{Canary: true},
				},
			},
			want: "test-12895109033131531539",
		},
		{
			name: "Test without canary",
			args: args{
				i: item{
					Name:      "test",
					Tags:      []string{"tag1", "tag2"},
					ExtraConf: configuration{Canary: false},
				},
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getName(tt.args.i); got != tt.want {
				t.Errorf("getName() = %v, want %v", got, tt.want)
			}
		})
	}
}

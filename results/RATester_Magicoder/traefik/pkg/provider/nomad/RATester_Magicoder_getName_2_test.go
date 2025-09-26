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
			name: "test_case_1",
			args: args{
				i: item{
					Name: "test_service",
					Tags: []string{"tag1", "tag2"},
					ExtraConf: configuration{
						Canary: false,
					},
				},
			},
			want: "test_service",
		},
		{
			name: "test_case_2",
			args: args{
				i: item{
					Name: "test_service",
					Tags: []string{"tag1", "tag2"},
					ExtraConf: configuration{
						Canary: true,
					},
				},
			},
			want: "test_service-1094404560771515913",
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

package consulcatalog

import (
	"fmt"
	"testing"
)

func TestGetName_3(t *testing.T) {
	type args struct {
		i itemData
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test getName with Canary false",
			args: args{
				i: itemData{
					Name: "test",
					ExtraConf: configuration{
						ConsulCatalog: specificConfiguration{
							Canary: false,
						},
					},
				},
			},
			want: "test",
		},
		{
			name: "Test getName with Canary true",
			args: args{
				i: itemData{
					Name: "test",
					Tags: []string{"tag1", "tag2", "tag3"},
					ExtraConf: configuration{
						ConsulCatalog: specificConfiguration{
							Canary: true,
						},
					},
				},
			},
			want: "test-12895109060342090268",
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

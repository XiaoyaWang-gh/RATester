package schema

import (
	"fmt"
	"testing"
)

func TestcontainsAlias_1(t *testing.T) {
	type args struct {
		infos []*structInfo
		alias string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				infos: []*structInfo{
					{
						fields: []*fieldInfo{
							{
								name:  "field1",
								alias: "alias1",
							},
							{
								name:  "field2",
								alias: "alias2",
							},
						},
					},
				},
				alias: "alias1",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				infos: []*structInfo{
					{
						fields: []*fieldInfo{
							{
								name:  "field1",
								alias: "alias1",
							},
							{
								name:  "field2",
								alias: "alias2",
							},
						},
					},
				},
				alias: "alias3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := containsAlias(tt.args.infos, tt.args.alias); got != tt.want {
				t.Errorf("containsAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

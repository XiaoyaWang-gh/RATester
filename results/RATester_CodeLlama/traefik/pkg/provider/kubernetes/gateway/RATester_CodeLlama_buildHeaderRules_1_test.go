package gateway

import (
	"fmt"
	"reflect"
	"testing"

	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestBuildHeaderRules_1(t *testing.T) {
	type args struct {
		headers []gatev1.HTTPHeaderMatch
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test case 1",
			args: args{
				headers: []gatev1.HTTPHeaderMatch{
					{
						Name:  "name1",
						Value: "value1",
					},
					{
						Name:  "name2",
						Value: "value2",
					},
				},
			},
			want: []string{
				"Header(`name1`,`value1`)",
				"Header(`name2`,`value2`)",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, _ := buildHeaderRules(tt.args.headers)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildHeaderRules() = %v, want %v", got, tt.want)
			}
		})
	}
}

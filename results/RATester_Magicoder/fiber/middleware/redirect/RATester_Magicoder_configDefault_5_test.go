package redirect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconfigDefault_5(t *testing.T) {
	type args struct {
		config []Config
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Test Case 1",
			args: args{
				config: []Config{
					{
						StatusCode: 301,
					},
				},
			},
			want: Config{
				StatusCode: 301,
			},
		},
		{
			name: "Test Case 2",
			args: args{
				config: []Config{
					{
						StatusCode: 0,
					},
				},
			},
			want: Config{
				StatusCode: ConfigDefault.StatusCode,
			},
		},
		{
			name: "Test Case 3",
			args: args{
				config: []Config{
					{
						StatusCode: 0,
					},
				},
			},
			want: Config{
				StatusCode: ConfigDefault.StatusCode,
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

			if got := configDefault(tt.args.config...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

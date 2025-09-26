package csrf

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfigDefault_13(t *testing.T) {
	type args struct {
		config []Config
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Test with empty config",
			args: args{
				config: []Config{},
			},
			want: ConfigDefault,
		},
		{
			name: "Test with custom config",
			args: args{
				config: []Config{
					{
						KeyLookup: "header:X-CSRF-TOKEN",
					},
				},
			},
			want: Config{
				KeyLookup: "header:X-CSRF-TOKEN",
				Extractor: FromHeader("X-CSRF-TOKEN"),
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

			got := configDefault(tt.args.config...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

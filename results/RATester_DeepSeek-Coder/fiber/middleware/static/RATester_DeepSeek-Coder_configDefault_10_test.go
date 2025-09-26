package static

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConfigDefault_10(t *testing.T) {
	type args struct {
		config []Config
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Test with no config",
			args: args{
				config: []Config{},
			},
			want: ConfigDefault,
		},
		{
			name: "Test with config",
			args: args{
				config: []Config{
					{
						IndexNames:    []string{"index.html", "custom.html"},
						CacheDuration: 20 * time.Second,
					},
				},
			},
			want: Config{
				IndexNames:    []string{"index.html", "custom.html"},
				CacheDuration: 20 * time.Second,
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

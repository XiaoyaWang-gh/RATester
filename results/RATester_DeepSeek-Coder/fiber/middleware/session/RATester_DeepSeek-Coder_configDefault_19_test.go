package session

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConfigDefault_19(t *testing.T) {
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
						Expiration: 1 * time.Hour,
						KeyLookup:  "header:session_id",
					},
				},
			},
			want: Config{
				Expiration:  1 * time.Hour,
				KeyLookup:   "header:session_id",
				source:      SourceHeader,
				sessionName: "session_id",
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

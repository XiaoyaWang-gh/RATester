package retry

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestconfigDefault_14(t *testing.T) {
	tests := []struct {
		name string
		args []Config
		want Config
	}{
		{
			name: "no args",
			args: []Config{},
			want: DefaultConfig,
		},
		{
			name: "with args",
			args: []Config{
				{
					InitialInterval: 2 * time.Second,
					MaxBackoffTime:  30 * time.Second,
					Multiplier:      1.5,
					MaxRetryCount:   15,
					currentInterval: 2 * time.Second,
				},
			},
			want: Config{
				InitialInterval: 2 * time.Second,
				MaxBackoffTime:  30 * time.Second,
				Multiplier:      1.5,
				MaxRetryCount:   15,
				currentInterval: 2 * time.Second,
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

			if got := configDefault(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

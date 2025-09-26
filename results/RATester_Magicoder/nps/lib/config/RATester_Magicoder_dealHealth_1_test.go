package config

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestdealHealth_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *file.Health
	}{
		{
			name: "Test case 1",
			args: args{s: "health_check_timeout=10,health_check_max_failed=5,health_check_interval=15,health_http_url=http://example.com,health_check_type=http,health_check_target=/health"},
			want: &file.Health{
				HealthCheckTimeout:  10,
				HealthMaxFail:       5,
				HealthCheckInterval: 15,
				HttpHealthUrl:       "http://example.com",
				HealthCheckType:     "http",
				HealthCheckTarget:   "/health",
			},
		},
		{
			name: "Test case 2",
			args: args{s: "health_check_timeout=20,health_check_max_failed=10,health_check_interval=20,health_http_url=http://example.org,health_check_type=tcp,health_check_target=127.0.0.1:8080"},
			want: &file.Health{
				HealthCheckTimeout:  20,
				HealthMaxFail:       10,
				HealthCheckInterval: 20,
				HttpHealthUrl:       "http://example.org",
				HealthCheckType:     "tcp",
				HealthCheckTarget:   "127.0.0.1:8080",
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

			if got := dealHealth(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dealHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

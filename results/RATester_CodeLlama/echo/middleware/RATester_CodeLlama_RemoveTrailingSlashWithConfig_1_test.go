package middleware

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestRemoveTrailingSlashWithConfig_1(t *testing.T) {
	type args struct {
		config TrailingSlashConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test case 1",
			args: args{
				config: TrailingSlashConfig{
					Skipper: func(c echo.Context) bool {
						return false
					},
				},
			},
		},
		{
			name: "test case 2",
			args: args{
				config: TrailingSlashConfig{
					Skipper: func(c echo.Context) bool {
						return true
					},
				},
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

			RemoveTrailingSlashWithConfig(tt.args.config)
		})
	}
}

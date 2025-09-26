package middleware

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

func TestToMiddleware_1(t *testing.T) {
	type args struct {
		config ContextTimeoutConfig
	}
	tests := []struct {
		name    string
		args    args
		want    echo.MiddlewareFunc
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				config: ContextTimeoutConfig{
					Timeout: 10 * time.Second,
				},
			},
			want: func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					if c.Request().Context().Done() != nil {
						return echo.ErrServiceUnavailable.WithInternal(c.Request().Context().Err())
					}
					return next(c)
				}
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				config: ContextTimeoutConfig{
					Timeout: 0,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.args.config.ToMiddleware()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMiddleware() got = %v, want %v", got, tt.want)
			}
		})
	}
}

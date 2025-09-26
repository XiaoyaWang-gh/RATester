package middleware

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestAddTarget_1(t *testing.T) {
	type args struct {
		target *ProxyTarget
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Add new target",
			args: args{
				target: &ProxyTarget{
					Name: "target1",
					URL:  &url.URL{},
					Meta: echo.Map{},
				},
			},
			want: true,
		},
		{
			name: "Add existing target",
			args: args{
				target: &ProxyTarget{
					Name: "target1",
					URL:  &url.URL{},
					Meta: echo.Map{},
				},
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

			b := &commonBalancer{
				targets: []*ProxyTarget{
					{
						Name: "target1",
						URL:  &url.URL{},
						Meta: echo.Map{},
					},
				},
			}
			if got := b.AddTarget(tt.args.target); got != tt.want {
				t.Errorf("commonBalancer.AddTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

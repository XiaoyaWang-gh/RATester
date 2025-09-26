package proxy

import (
	"fmt"
	"reflect"
	"testing"

	"golang.org/x/time/rate"
)

func TestGetLimiter_1(t *testing.T) {
	type fields struct {
		limiter *rate.Limiter
	}
	tests := []struct {
		name   string
		fields fields
		want   *rate.Limiter
	}{
		{
			name: "TestGetLimiter",
			fields: fields{
				limiter: rate.NewLimiter(1, 1),
			},
			want: rate.NewLimiter(1, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			pxy := &BaseProxy{
				limiter: tt.fields.limiter,
			}
			if got := pxy.GetLimiter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseProxy.GetLimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

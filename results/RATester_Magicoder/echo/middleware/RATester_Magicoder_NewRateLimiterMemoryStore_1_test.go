package middleware

import (
	"fmt"
	"reflect"
	"testing"

	"golang.org/x/time/rate"
)

func TestNewRateLimiterMemoryStore_1(t *testing.T) {
	type args struct {
		rate rate.Limit
	}
	tests := []struct {
		name string
		args args
		want *RateLimiterMemoryStore
	}{
		{
			name: "Test case 1",
			args: args{
				rate: 10,
			},
			want: &RateLimiterMemoryStore{
				visitors: make(map[string]*Visitor),
				rate:     10,
			},
		},
		{
			name: "Test case 2",
			args: args{
				rate: 20,
			},
			want: &RateLimiterMemoryStore{
				visitors: make(map[string]*Visitor),
				rate:     20,
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

			if got := NewRateLimiterMemoryStore(tt.args.rate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRateLimiterMemoryStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

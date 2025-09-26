package requestdecorator

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

func TestCNAMEFlatten_1(t *testing.T) {
	ctx := context.Background()
	resolver := &Resolver{
		CnameFlattening: true,
		ResolvConfig:    "test_config",
		ResolvDepth:     2,
		cache:           cache.New(30*time.Minute, 5*time.Minute),
	}

	tests := []struct {
		name string
		host string
		want string
	}{
		{
			name: "Test case 1",
			host: "example.com",
			want: "example.com",
		},
		{
			name: "Test case 2",
			host: "example2.com",
			want: "example2.com",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := resolver.CNAMEFlatten(ctx, tt.host); got != tt.want {
				t.Errorf("CNAMEFlatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

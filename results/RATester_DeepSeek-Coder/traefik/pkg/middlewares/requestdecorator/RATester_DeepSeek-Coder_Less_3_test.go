package requestdecorator

import (
	"fmt"
	"testing"
	"time"
)

func TestLess_3(t *testing.T) {
	tests := []struct {
		name string
		a    byTTL
		i    int
		j    int
		want bool
	}{
		{
			name: "Test case 1",
			a:    byTTL{&cnameResolv{TTL: 1 * time.Second, Record: "test1"}, &cnameResolv{TTL: 2 * time.Second, Record: "test2"}},
			i:    0,
			j:    1,
			want: true,
		},
		{
			name: "Test case 2",
			a:    byTTL{&cnameResolv{TTL: 1 * time.Second, Record: "test1"}, &cnameResolv{TTL: 2 * time.Second, Record: "test2"}},
			i:    1,
			j:    0,
			want: false,
		},
		{
			name: "Test case 3",
			a:    byTTL{&cnameResolv{TTL: 1 * time.Second, Record: "test1"}, &cnameResolv{TTL: 1 * time.Second, Record: "test2"}},
			i:    0,
			j:    1,
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

			if got := tt.a.Less(tt.i, tt.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

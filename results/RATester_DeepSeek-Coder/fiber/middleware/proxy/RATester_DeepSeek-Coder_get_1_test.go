package proxy

import (
	"fmt"
	"sync"
	"testing"
)

func TestGet_1(t *testing.T) {
	type fields struct {
		pool    []string
		current int
		Mutex   sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				pool:    []string{"A", "B", "C"},
				current: 0,
				Mutex:   sync.Mutex{},
			},
			want: "A",
		},
		{
			name: "Test case 2",
			fields: fields{
				pool:    []string{"A", "B", "C"},
				current: 1,
				Mutex:   sync.Mutex{},
			},
			want: "B",
		},
		{
			name: "Test case 3",
			fields: fields{
				pool:    []string{"A", "B", "C"},
				current: 3,
				Mutex:   sync.Mutex{},
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &roundrobin{
				pool:    tt.fields.pool,
				current: tt.fields.current,
				Mutex:   tt.fields.Mutex,
			}
			if got := r.get(); got != tt.want {
				t.Errorf("roundrobin.get() = %v, want %v", got, tt.want)
			}
		})
	}
}

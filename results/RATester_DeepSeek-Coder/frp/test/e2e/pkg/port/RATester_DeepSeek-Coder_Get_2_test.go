package port

import (
	"fmt"
	"sync"
	"testing"

	"k8s.io/apimachinery/pkg/util/sets"
)

func TestGet_2(t *testing.T) {
	type fields struct {
		reserved sets.Set[int]
		used     sets.Set[int]
		mu       sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			pa := &Allocator{
				reserved: tt.fields.reserved,
				used:     tt.fields.used,
				mu:       tt.fields.mu,
			}
			if got := pa.Get(); got != tt.want {
				t.Errorf("Allocator.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

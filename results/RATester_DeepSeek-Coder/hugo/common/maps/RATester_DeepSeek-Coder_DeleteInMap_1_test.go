package maps

import (
	"fmt"
	"sync"
	"testing"
)

func TestDeleteInMap_1(t *testing.T) {
	type fields struct {
		values map[string]any
		mu     sync.RWMutex
	}
	type args struct {
		key    string
		mapKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
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

			c := &Scratch{
				values: tt.fields.values,
				mu:     tt.fields.mu,
			}
			if got := c.DeleteInMap(tt.args.key, tt.args.mapKey); got != tt.want {
				t.Errorf("Scratch.DeleteInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

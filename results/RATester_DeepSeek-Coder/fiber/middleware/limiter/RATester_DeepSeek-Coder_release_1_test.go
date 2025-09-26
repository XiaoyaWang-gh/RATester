package limiter

import (
	"fmt"
	"sync"
	"testing"
)

func TestRelease_1(t *testing.T) {
	type args struct {
		e *item
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				e: &item{
					currHits: 10,
					prevHits: 5,
					exp:      100,
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				e: &item{
					currHits: 0,
					prevHits: 0,
					exp:      0,
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

			m := &manager{
				pool: sync.Pool{
					New: func() interface{} {
						return &item{}
					},
				},
			}
			m.release(tt.args.e)
			if tt.args.e.currHits != 0 || tt.args.e.prevHits != 0 || tt.args.e.exp != 0 {
				t.Errorf("Expected currHits, prevHits and exp to be 0, but got %d, %d, %d", tt.args.e.currHits, tt.args.e.prevHits, tt.args.e.exp)
			}
		})
	}
}

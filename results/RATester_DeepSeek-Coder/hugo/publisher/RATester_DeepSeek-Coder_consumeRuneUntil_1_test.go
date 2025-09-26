package publisher

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConsumeRuneUntil_1(t *testing.T) {
	tests := []struct {
		name string
		rune rune
		want htmlCollectorStateFunc
	}{
		{
			name: "Test consumeRuneUntil with condition that returns true",
			rune: 'a',
			want: func(*htmlElementsCollectorWriter) htmlCollectorStateFunc {
				return nil
			},
		},
		{
			name: "Test consumeRuneUntil with condition that returns false",
			rune: 'b',
			want: func(*htmlElementsCollectorWriter) htmlCollectorStateFunc {
				return nil
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

			w := &htmlElementsCollectorWriter{
				r: tt.rune,
			}
			if got := w.consumeRuneUntil(func(r rune) bool {
				return r == tt.rune
			}, tt.want); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("consumeRuneUntil() = %v, want %v", got, tt.want)
			}
		})
	}
}

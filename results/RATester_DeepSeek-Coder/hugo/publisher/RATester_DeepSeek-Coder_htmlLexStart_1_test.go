package publisher

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHtmlLexStart_1(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want htmlCollectorStateFunc
	}{
		{"Test htmlLexStart with '<'", '<', htmlLexElementStart},
		{"Test htmlLexStart with '>'", '>', htmlLexStart},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := &htmlElementsCollectorWriter{
				r: tt.r,
			}
			if got := htmlLexStart(w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("htmlLexStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

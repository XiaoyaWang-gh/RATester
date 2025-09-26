package publisher

import (
	"fmt"
	"reflect"
	"testing"
)

func TesthtmlLexStart_1(t *testing.T) {
	tests := []struct {
		name string
		w    *htmlElementsCollectorWriter
		want htmlCollectorStateFunc
	}{
		{
			name: "Test case 1",
			w: &htmlElementsCollectorWriter{
				r: '<',
			},
			want: htmlLexElementStart,
		},
		{
			name: "Test case 2",
			w: &htmlElementsCollectorWriter{
				r: '>',
			},
			want: htmlLexStart,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := htmlLexStart(tt.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("htmlLexStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

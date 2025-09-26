package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEmit_3(t *testing.T) {
	tests := []struct {
		name string
		l    *pageLexer
		t    ItemType
		want []Item
	}{
		{
			name: "test emit with tText",
			l: &pageLexer{
				input: []byte("  hello"),
				pos:   5,
				start: 2,
				items: []Item{},
			},
			t: tText,
			want: []Item{
				{Type: tText, low: 2, high: 5},
			},
		},
		{
			name: "test emit with tIndentation",
			l: &pageLexer{
				input: []byte("  hello"),
				pos:   5,
				start: 2,
				items: []Item{},
			},
			t: tIndentation,
			want: []Item{
				{Type: tIndentation, low: 2, high: 5},
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

			tt.l.emit(tt.t)
			if !reflect.DeepEqual(tt.l.items, tt.want) {
				t.Errorf("got %v, want %v", tt.l.items, tt.want)
			}
		})
	}
}

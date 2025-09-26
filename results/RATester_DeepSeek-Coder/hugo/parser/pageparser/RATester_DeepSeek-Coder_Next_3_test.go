package pageparser

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNext_3(t *testing.T) {
	type fields struct {
		items   Items
		lastPos int
	}
	tests := []struct {
		name   string
		fields fields
		want   Item
	}{
		{
			name: "TestNext_0",
			fields: fields{
				items:   Items{Item{Type: 0, Err: nil, low: 0, high: 1}},
				lastPos: 0,
			},
			want: Item{Type: 0, Err: nil, low: 0, high: 1},
		},
		{
			name: "TestNext_1",
			fields: fields{
				items:   Items{Item{Type: 1, Err: errors.New("error"), low: 1, high: 2}},
				lastPos: 0,
			},
			want: Item{Type: 1, Err: errors.New("error"), low: 1, high: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			iter := &Iterator{
				items:   tt.fields.items,
				lastPos: tt.fields.lastPos,
			}
			if got := iter.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Iterator.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

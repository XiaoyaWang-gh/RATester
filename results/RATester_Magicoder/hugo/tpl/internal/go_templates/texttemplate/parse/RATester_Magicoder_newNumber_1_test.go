package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewNumber_1(t *testing.T) {
	tests := []struct {
		name    string
		pos     Pos
		text    string
		typ     itemType
		want    *NumberNode
		wantErr bool
	}{
		{
			name: "test char constant",
			pos:  1,
			text: "'a'",
			typ:  itemCharConstant,
			want: &NumberNode{
				NodeType: NodeNumber,
				Pos:      1,
				Text:     "'a'",
				IsInt:    true,
				Uint64:   97,
				IsUint:   true,
				Float64:  97,
				IsFloat:  true,
				Int64:    97,
			},
			wantErr: false,
		},
		{
			name: "test complex",
			pos:  1,
			text: "1+2i",
			typ:  itemComplex,
			want: &NumberNode{
				NodeType:   NodeNumber,
				Pos:        1,
				Text:       "1+2i",
				IsComplex:  true,
				Complex128: complex(1, 2),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tr := &Tree{}
			got, err := tr.newNumber(tt.pos, tt.text, tt.typ)
			if (err != nil) != tt.wantErr {
				t.Errorf("newNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

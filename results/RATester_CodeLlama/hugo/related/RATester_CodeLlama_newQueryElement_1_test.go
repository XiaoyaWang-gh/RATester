package related

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewQueryElement_1(t *testing.T) {
	type args struct {
		index    string
		keywords []Keyword
	}
	tests := []struct {
		name string
		args args
		want queryElement
	}{
		{
			name: "test1",
			args: args{
				index:    "test1",
				keywords: []Keyword{},
			},
			want: queryElement{
				Index:    "test1",
				Keywords: []Keyword{},
			},
		},
		{
			name: "test2",
			args: args{
				index:    "test2",
				keywords: []Keyword{},
			},
			want: queryElement{
				Index:    "test2",
				Keywords: []Keyword{},
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

			if got := newQueryElement(tt.args.index, tt.args.keywords...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueryElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

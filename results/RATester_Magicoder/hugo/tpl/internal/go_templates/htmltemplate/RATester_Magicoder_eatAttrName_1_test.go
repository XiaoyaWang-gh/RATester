package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TesteatAttrName_1(t *testing.T) {
	tests := []struct {
		name  string
		s     []byte
		i     int
		want  int
		want1 *Error
	}{
		{
			name:  "Test case 1",
			s:     []byte("<tag attr='value'>"),
			i:     5,
			want:  14,
			want1: nil,
		},
		{
			name:  "Test case 2",
			s:     []byte("<tag attr=\"value\""),
			i:     5,
			want:  -1,
			want1: &Error{ErrorCode: ErrBadHTML},
		},
		{
			name:  "Test case 3",
			s:     []byte("<tag attr=value>"),
			i:     5,
			want:  13,
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := eatAttrName(tt.s, tt.i)
			if got != tt.want {
				t.Errorf("eatAttrName() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("eatAttrName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

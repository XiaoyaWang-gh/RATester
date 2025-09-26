package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLex_2(t *testing.T) {
	type fields struct {
		l           *pageLexer
		skipAll     bool
		handlers    []*sectionHandler
		skipIndexes []int
	}
	type args struct {
		origin stateFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   stateFunc
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

			s := &sectionHandlers{
				l:           tt.fields.l,
				skipAll:     tt.fields.skipAll,
				handlers:    tt.fields.handlers,
				skipIndexes: tt.fields.skipIndexes,
			}
			if got := s.lex(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sectionHandlers.lex() = %v, want %v", got, tt.want)
			}
		})
	}
}

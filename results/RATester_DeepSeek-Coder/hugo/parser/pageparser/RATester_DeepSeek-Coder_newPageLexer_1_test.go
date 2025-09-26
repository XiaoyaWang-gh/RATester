package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewPageLexer_1(t *testing.T) {
	type args struct {
		input      []byte
		stateStart stateFunc
		cfg        Config
	}
	tests := []struct {
		name string
		args args
		want *pageLexer
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

			if got := newPageLexer(tt.args.input, tt.args.stateStart, tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPageLexer() = %v, want %v", got, tt.want)
			}
		})
	}
}

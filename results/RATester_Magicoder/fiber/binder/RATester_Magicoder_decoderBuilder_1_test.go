package binder

import (
	"fmt"
	"reflect"
	"testing"
)

func TestdecoderBuilder_1(t *testing.T) {
	type args struct {
		parserConfig ParserConfig
	}
	tests := []struct {
		name string
		args args
		want any
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

			if got := decoderBuilder(tt.args.parserConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decoderBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

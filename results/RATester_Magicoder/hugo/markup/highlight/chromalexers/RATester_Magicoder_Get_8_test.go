package chromalexers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alecthomas/chroma/v2"
)

func TestGet_8(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want chroma.Lexer
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

			if got := Get(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

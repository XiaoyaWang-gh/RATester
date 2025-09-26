package publisher

import (
	"fmt"
	"reflect"
	"testing"
)

func TesthtmlLexToEndOfComment_1(t *testing.T) {
	tests := []struct {
		name string
		w    *htmlElementsCollectorWriter
		want htmlCollectorStateFunc
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

			if got := htmlLexToEndOfComment(tt.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("htmlLexToEndOfComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

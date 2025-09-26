package publisher

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/transform"
)

func TestcreateTransformerChain_1(t *testing.T) {
	type args struct {
		f Descriptor
	}
	tests := []struct {
		name string
		p    DestinationPublisher
		args args
		want transform.Chain
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

			if got := tt.p.createTransformerChain(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DestinationPublisher.createTransformerChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

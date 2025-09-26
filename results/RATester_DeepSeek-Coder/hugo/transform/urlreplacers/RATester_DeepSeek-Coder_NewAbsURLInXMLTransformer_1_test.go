package urlreplacers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/transform"
)

func TestNewAbsURLInXMLTransformer_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want transform.Transformer
	}{
		{
			name: "Test case 1",
			args: args{
				path: "test_path",
			},
			want: func(ft transform.FromTo) error {
				ar.replaceInXML("test_path", ft)
				return nil
			},
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewAbsURLInXMLTransformer(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAbsURLInXMLTransformer() = %v, want %v", got, tt.want)
			}
		})
	}
}

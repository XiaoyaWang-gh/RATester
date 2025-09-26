package segments

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestDecodeSegments_1(t *testing.T) {
	type args struct {
		in map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    *config.ConfigNamespace[map[string]SegmentConfig, Segments]
		wantErr bool
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

			got, err := DecodeSegments(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeSegments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}

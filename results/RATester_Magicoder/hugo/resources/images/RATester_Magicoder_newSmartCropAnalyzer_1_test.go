package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
	"github.com/muesli/smartcrop"
)

func TestnewSmartCropAnalyzer_1(t *testing.T) {
	type args struct {
		filter gift.Resampling
	}
	tests := []struct {
		name string
		p    *ImageProcessor
		args args
		want smartcrop.Analyzer
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

			if got := tt.p.newSmartCropAnalyzer(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSmartCropAnalyzer() = %v, want %v", got, tt.want)
			}
		})
	}
}

package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestParse_2(t *testing.T) {
	type args struct {
		ctx converter.RenderContext
	}
	tests := []struct {
		name    string
		c       *goldmarkConverter
		args    args
		want    converter.ResultParse
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

			got, err := tt.c.Parse(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("goldmarkConverter.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goldmarkConverter.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

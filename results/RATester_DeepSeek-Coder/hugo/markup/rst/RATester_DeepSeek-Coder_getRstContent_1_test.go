package rst

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestGetRstContent_1(t *testing.T) {
	type args struct {
		src []byte
		ctx converter.DocumentContext
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
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

			c := &rstConverter{}
			got, err := c.getRstContent(tt.args.src, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRstContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRstContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

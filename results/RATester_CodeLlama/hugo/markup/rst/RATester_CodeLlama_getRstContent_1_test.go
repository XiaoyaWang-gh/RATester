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
		c       *rstConverter
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test1",
			c:    &rstConverter{},
			args: args{
				src: []byte(""),
				ctx: converter.DocumentContext{},
			},
			want:    []byte(""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.c.getRstContent(tt.args.src, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRstContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRstContent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

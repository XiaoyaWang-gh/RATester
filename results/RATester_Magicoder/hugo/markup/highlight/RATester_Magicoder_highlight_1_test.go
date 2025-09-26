package highlight

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/hugio"
	"github.com/gohugoio/hugo/markup/internal/attributes"
)

func Testhighlight_1(t *testing.T) {
	type args struct {
		fw    hugio.FlexiWriter
		code  string
		lang  string
		attrs []attributes.Attribute
		cfg   Config
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
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

			got, got1, err := highlight(tt.args.fw, tt.args.code, tt.args.lang, tt.args.attrs, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("highlight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("highlight() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("highlight() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

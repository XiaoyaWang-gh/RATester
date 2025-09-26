package fmt

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
)

func TestWarnf_1(t *testing.T) {
	type fields struct {
		logger loggers.Logger
	}
	type args struct {
		format string
		args   []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
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

			ns := &Namespace{
				logger: tt.fields.logger,
			}
			if got := ns.Warnf(tt.args.format, tt.args.args...); got != tt.want {
				t.Errorf("Namespace.Warnf() = %v, want %v", got, tt.want)
			}
		})
	}
}

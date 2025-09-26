package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types"
)

func TestValue_2(t *testing.T) {
	type args struct {
		l types.LowHigh[string]
	}
	tests := []struct {
		name string
		s    HtmlSummary
		args args
		want string
	}{
		{
			name: "test1",
			s:    HtmlSummary{},
			args: args{l: types.LowHigh[string]{Low: 0, High: 10}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.Value(tt.args.l); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

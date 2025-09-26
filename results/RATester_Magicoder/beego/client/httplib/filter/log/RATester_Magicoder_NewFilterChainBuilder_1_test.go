package log

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/logs"
)

func TestNewFilterChainBuilder_1(t *testing.T) {
	tests := []struct {
		name string
		want *FilterChainBuilder
	}{
		{
			name: "Test case 1",
			want: &FilterChainBuilder{
				printableContentTypes: defaultprintableContentTypes,
				log:                   logs.Debug,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewFilterChainBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterChainBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

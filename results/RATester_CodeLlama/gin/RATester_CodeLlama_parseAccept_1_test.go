package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseAccept_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "test1",
			args: "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			want: []string{"text/html", "application/xhtml+xml", "application/xml", "image/webp", "image/apng", "*/*", "application/signed-exchange;v=b3;q=0.9"},
		},
		{
			name: "test2",
			args: "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			want: []string{"text/html", "application/xhtml+xml", "application/xml", "image/webp", "image/apng", "*/*", "application/signed-exchange;v=b3;q=0.9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := parseAccept(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseAccept() = %v, want %v", got, tt.want)
			}
		})
	}
}

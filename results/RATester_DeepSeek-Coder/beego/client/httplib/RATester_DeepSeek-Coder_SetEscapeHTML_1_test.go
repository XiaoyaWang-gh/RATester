package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetEscapeHTML_1(t *testing.T) {
	type args struct {
		isEscape bool
	}
	tests := []struct {
		name string
		args args
		want *BeegoHTTPRequest
	}{
		{
			name: "TestSetEscapeHTML_True",
			args: args{
				isEscape: true,
			},
			want: &BeegoHTTPRequest{
				setting: BeegoHTTPSettings{
					EscapeHTML: true,
				},
			},
		},
		{
			name: "TestSetEscapeHTML_False",
			args: args{
				isEscape: false,
			},
			want: &BeegoHTTPRequest{
				setting: BeegoHTTPSettings{
					EscapeHTML: false,
				},
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

			b := &BeegoHTTPRequest{
				setting: BeegoHTTPSettings{
					EscapeHTML: true,
				},
			}
			if got := b.SetEscapeHTML(tt.args.isEscape); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeegoHTTPRequest.SetEscapeHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}

package mock

import (
	"fmt"
	"net/textproto"
	"reflect"
	"testing"
)

func TestWithHeader_2(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want simpleConditionOption
	}{
		{
			name: "Test case 1",
			args: args{
				key:   "Content-Type",
				value: "application/json",
			},
			want: func(sc *SimpleCondition) {
				sc.header[textproto.CanonicalMIMEHeaderKey("Content-Type")] = "application/json"
			},
		},
		{
			name: "Test case 2",
			args: args{
				key:   "Accept",
				value: "*/*",
			},
			want: func(sc *SimpleCondition) {
				sc.header[textproto.CanonicalMIMEHeaderKey("Accept")] = "*/*"
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

			if got := WithHeader(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

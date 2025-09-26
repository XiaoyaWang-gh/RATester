package streamserver

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithRespContent_1(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "TestWithRespContent",
			args: args{
				content: []byte("test content"),
			},
			want: []byte("test content"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := WithRespContent(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRespContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

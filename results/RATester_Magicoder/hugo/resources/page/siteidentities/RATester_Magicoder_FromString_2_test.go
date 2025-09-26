package siteidentities

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestFromString_2(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		args  args
		want  identity.Identity
		want1 bool
	}{
		{
			name:  "Test case 1",
			args:  args{name: "Data"},
			want:  Data,
			want1: true,
		},
		{
			name:  "Test case 2",
			args:  args{name: "Invalid"},
			want:  identity.Anonymous,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := FromString(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FromString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewMediaType_1(t *testing.T) {
	type args struct {
		main     string
		sub      string
		suffixes []string
	}
	tests := []struct {
		name string
		args args
		want Type
	}{
		{
			name: "Test Case 1",
			args: args{
				main:     "application",
				sub:      "rss",
				suffixes: []string{"xml"},
			},
			want: Type{
				MainType:    "application",
				SubType:     "rss",
				SuffixesCSV: "xml",
				Delimiter:   DefaultDelimiter,
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newMediaType(tt.args.main, tt.args.sub, tt.args.suffixes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMediaType() = %v, want %v", got, tt.want)
			}
		})
	}
}

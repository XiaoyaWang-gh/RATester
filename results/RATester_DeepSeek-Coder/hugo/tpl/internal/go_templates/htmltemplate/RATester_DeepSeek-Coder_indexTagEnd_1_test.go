package template

import (
	"fmt"
	"testing"
)

func TestIndexTagEnd_1(t *testing.T) {
	type args struct {
		s   []byte
		tag []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				s:   []byte("<tag>content</tag>"),
				tag: []byte("tag"),
			},
			want: 0,
		},
		{
			name: "Test Case 2",
			args: args{
				s:   []byte("<tag>content</tag><tag>content</tag>"),
				tag: []byte("tag"),
			},
			want: 14,
		},
		{
			name: "Test Case 3",
			args: args{
				s:   []byte("<tag>content</tag><tag>content</tag>"),
				tag: []byte("notag"),
			},
			want: -1,
		},
		{
			name: "Test Case 4",
			args: args{
				s:   []byte("<tag>content<tag>content</tag>"),
				tag: []byte("tag"),
			},
			want: -1,
		},
		{
			name: "Test Case 5",
			args: args{
				s:   []byte("<tag>content</tag>"),
				tag: []byte("TAG"),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := indexTagEnd(tt.args.s, tt.args.tag); got != tt.want {
				t.Errorf("indexTagEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

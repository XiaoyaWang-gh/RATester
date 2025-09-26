package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func Testsource_1(t *testing.T) {
	type args struct {
		lines [][]byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test Case 1",
			args: args{
				lines: [][]byte{[]byte("line1"), []byte("line2"), []byte("line3")},
				n:     1,
			},
			want: []byte("line2"),
		},
		{
			name: "Test Case 2",
			args: args{
				lines: [][]byte{[]byte("line1"), []byte("line2"), []byte("line3")},
				n:     3,
			},
			want: []byte("line3"),
		},
		{
			name: "Test Case 3",
			args: args{
				lines: [][]byte{[]byte("line1"), []byte("line2"), []byte("line3")},
				n:     0,
			},
			want: []byte("line1"),
		},
		{
			name: "Test Case 4",
			args: args{
				lines: [][]byte{[]byte("line1"), []byte("line2"), []byte("line3")},
				n:     4,
			},
			want: nil,
		},
		{
			name: "Test Case 5",
			args: args{
				lines: [][]byte{[]byte("line1"), []byte("line2"), []byte("line3")},
				n:     -1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := source(tt.args.lines, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("source() = %v, want %v", got, tt.want)
			}
		})
	}
}

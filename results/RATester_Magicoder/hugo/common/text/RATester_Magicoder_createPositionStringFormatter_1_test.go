package text

import (
	"fmt"
	"testing"
)

func TestcreatePositionStringFormatter_1(t *testing.T) {
	type args struct {
		formatStr string
		pos       Position
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				formatStr: ":file::line::col",
				pos: Position{
					Filename:     "test.go",
					LineNumber:   10,
					ColumnNumber: 5,
				},
			},
			want: "test.go:10:5",
		},
		{
			name: "Test case 2",
			args: args{
				formatStr: ":file::line::col",
				pos: Position{
					Filename:     "test.go",
					LineNumber:   15,
					ColumnNumber: 10,
				},
			},
			want: "test.go:15:10",
		},
		{
			name: "Test case 3",
			args: args{
				formatStr: ":file::line::col",
				pos: Position{
					Filename:     "main.go",
					LineNumber:   20,
					ColumnNumber: 15,
				},
			},
			want: "main.go:20:15",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := createPositionStringFormatter(tt.args.formatStr)
			if got := f(tt.args.pos); got != tt.want {
				t.Errorf("createPositionStringFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

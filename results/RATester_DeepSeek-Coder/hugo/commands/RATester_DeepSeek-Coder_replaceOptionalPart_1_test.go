package commands

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReplaceOptionalPart_1(t *testing.T) {
	type args struct {
		buffer   *bytes.Buffer
		partName string
		part     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with empty part",
			args: args{
				buffer:   &bytes.Buffer{},
				partName: "test",
				part:     "",
			},
			want: "",
		},
		{
			name: "Test with non-empty part",
			args: args{
				buffer:   &bytes.Buffer{},
				partName: "test",
				part:     "test_part",
			},
			want: "test=\"test_part\" ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &importCommand{}
			c.replaceOptionalPart(tt.args.buffer, tt.args.partName, tt.args.part)
			got := tt.args.buffer.String()
			if got != tt.want {
				t.Errorf("replaceOptionalPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

package logs

import (
	"fmt"
	"testing"
)

func TestSetFormatter_4(t *testing.T) {
	type args struct {
		f LogFormatter
	}
	tests := []struct {
		name string
		c    *connWriter
		args args
		want LogFormatter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &connWriter{}
			c.SetFormatter(tt.args.f)
			if c.formatter != tt.want {
				t.Errorf("connWriter.SetFormatter() = %v, want %v", c.formatter, tt.want)
			}
		})
	}
}

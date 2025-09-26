package logs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetFormatter_5(t *testing.T) {
	type args struct {
		f LogFormatter
	}
	tests := []struct {
		name string
		s    *SMTPWriter
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

			s := &SMTPWriter{}
			s.SetFormatter(tt.args.f)
			if !reflect.DeepEqual(s.formatter, tt.want) {
				t.Errorf("SMTPWriter.SetFormatter() = %v, want %v", s.formatter, tt.want)
			}
		})
	}
}

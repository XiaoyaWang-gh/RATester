package logs

import (
	"fmt"
	"testing"
)

func TestFormat_4(t *testing.T) {
	type args struct {
		lm *LogMsg
	}
	tests := []struct {
		name string
		s    *SLACKWriter
		args args
		want string
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

			if got := tt.s.Format(tt.args.lm); got != tt.want {
				t.Errorf("SLACKWriter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

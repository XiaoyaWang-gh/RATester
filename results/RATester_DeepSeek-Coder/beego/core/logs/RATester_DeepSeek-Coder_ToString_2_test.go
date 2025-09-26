package logs

import (
	"fmt"
	"testing"
)

func TestToString_2(t *testing.T) {
	type args struct {
		lm *LogMsg
	}
	tests := []struct {
		name string
		p    *PatternLogFormatter
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

			if got := tt.p.ToString(tt.args.lm); got != tt.want {
				t.Errorf("PatternLogFormatter.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

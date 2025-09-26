package fmt

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func Testlogmf_1(t *testing.T) {
	type args struct {
		l      logg.LevelLogger
		m      any
		format string
		args   []any
	}
	tests := []struct {
		name string
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

			ns := &Namespace{}
			if got := ns.logmf(tt.args.l, tt.args.m, tt.args.format, tt.args.args...); got != tt.want {
				t.Errorf("logmf() = %v, want %v", got, tt.want)
			}
		})
	}
}

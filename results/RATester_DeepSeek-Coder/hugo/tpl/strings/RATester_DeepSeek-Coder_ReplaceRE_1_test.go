package strings

import (
	"fmt"
	"testing"
)

func TestReplaceRE_1(t *testing.T) {
	type args struct {
		pattern any
		repl    any
		s       any
		n       []any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
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
			got, err := ns.ReplaceRE(tt.args.pattern, tt.args.repl, tt.args.s, tt.args.n...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.ReplaceRE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.ReplaceRE() = %v, want %v", got, tt.want)
			}
		})
	}
}

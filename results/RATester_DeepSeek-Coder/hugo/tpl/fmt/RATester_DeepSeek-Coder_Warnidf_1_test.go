package fmt

import (
	"fmt"
	"testing"
)

func TestWarnidf_1(t *testing.T) {
	type args struct {
		id     string
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
			if got := ns.Warnidf(tt.args.id, tt.args.format, tt.args.args...); got != tt.want {
				t.Errorf("Namespace.Warnidf() = %v, want %v", got, tt.want)
			}
		})
	}
}

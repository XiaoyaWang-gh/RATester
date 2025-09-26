package compare

import (
	"fmt"
	"testing"
)

func TestCompareTwoUints_1(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
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
			got, got1 := ns.compareTwoUints(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("Namespace.compareTwoUints() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Namespace.compareTwoUints() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

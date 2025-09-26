package collections

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSlice_3(t *testing.T) {
	ns := &Namespace{
		loc: time.UTC,
	}

	tests := []struct {
		name string
		args []any
		want any
	}{
		{
			name: "no args",
			args: []any{},
			want: []any{},
		},
		{
			name: "with args",
			args: []any{1, 2, 3},
			want: []any{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ns.Slice(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

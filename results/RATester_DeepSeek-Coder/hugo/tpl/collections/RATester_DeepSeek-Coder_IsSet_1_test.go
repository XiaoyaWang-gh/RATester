package collections

import (
	"fmt"
	"testing"
	"time"

	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/tpl/compare"
)

func TestIsSet_1(t *testing.T) {
	ns := &Namespace{
		loc:      time.UTC,
		sortComp: &compare.Namespace{},
		deps:     &deps.Deps{},
	}

	tests := []struct {
		name    string
		ns      *Namespace
		arg1    any
		arg2    any
		want    bool
		wantErr bool
	}{
		{
			name: "Test with array",
			ns:   ns,
			arg1: []int{1, 2, 3},
			arg2: 1,
			want: true,
		},
		{
			name: "Test with slice",
			ns:   ns,
			arg1: []int{1, 2, 3},
			arg2: 1,
			want: true,
		},
		{
			name: "Test with map",
			ns:   ns,
			arg1: map[string]int{"a": 1, "b": 2, "c": 3},
			arg2: "a",
			want: true,
		},
		{
			name: "Test with unsupported type",
			ns:   ns,
			arg1: "test",
			arg2: "a",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.ns.IsSet(tt.arg1, tt.arg2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.IsSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

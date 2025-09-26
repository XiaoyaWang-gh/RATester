package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestcanBeNil_2(t *testing.T) {
	tests := []struct {
		name string
		typ  reflect.Type
		want bool
	}{
		{
			name: "Chan",
			typ:  reflect.ChanOf(reflect.SendDir, reflect.TypeOf(0)),
			want: true,
		},
		{
			name: "Func",
			typ:  reflect.FuncOf([]reflect.Type{}, []reflect.Type{}, false),
			want: true,
		},
		{
			name: "Interface",
			typ:  reflect.TypeOf((*interface{})(nil)).Elem(),
			want: true,
		},
		{
			name: "Map",
			typ:  reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0)),
			want: true,
		},
		{
			name: "Ptr",
			typ:  reflect.PtrTo(reflect.TypeOf(0)),
			want: true,
		},
		{
			name: "Slice",
			typ:  reflect.SliceOf(reflect.TypeOf(0)),
			want: true,
		},
		{
			name: "NotNil",
			typ:  reflect.TypeOf(0),
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

			if got := canBeNil(tt.typ); got != tt.want {
				t.Errorf("canBeNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

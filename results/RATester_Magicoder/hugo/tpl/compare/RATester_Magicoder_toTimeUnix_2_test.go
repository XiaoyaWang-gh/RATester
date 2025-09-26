package compare

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TesttoTimeUnix_2(t *testing.T) {
	ns := &Namespace{
		loc: time.Local,
	}

	tests := []struct {
		name string
		v    reflect.Value
		want int64
	}{
		{
			name: "valid time.Time",
			v:    reflect.ValueOf(time.Now()),
			want: time.Now().Unix(),
		},
		{
			name: "invalid type",
			v:    reflect.ValueOf("not a time.Time"),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ns.toTimeUnix(tt.v); got != tt.want {
				t.Errorf("toTimeUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

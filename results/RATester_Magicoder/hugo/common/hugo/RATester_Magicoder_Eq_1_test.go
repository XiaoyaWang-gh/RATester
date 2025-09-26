package hugo

import (
	"fmt"
	"testing"
)

func TestEq_1(t *testing.T) {
	tests := []struct {
		name string
		h    VersionString
		arg  any
		want bool
	}{
		{
			name: "Equal",
			h:    "1.2.3",
			arg:  "1.2.3",
			want: true,
		},
		{
			name: "NotEqual",
			h:    "1.2.3",
			arg:  "1.2.4",
			want: false,
		},
		{
			name: "NotString",
			h:    "1.2.3",
			arg:  123,
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

			if got := tt.h.Eq(tt.arg); got != tt.want {
				t.Errorf("Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

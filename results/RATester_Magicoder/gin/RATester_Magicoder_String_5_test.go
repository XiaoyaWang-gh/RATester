package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestString_5(t *testing.T) {
	tests := []struct {
		name string
		a    errorMsgs
		want string
	}{
		{
			name: "empty",
			a:    errorMsgs{},
			want: "",
		},
		{
			name: "single",
			a: errorMsgs{
				{
					Err:  errors.New("error 1"),
					Meta: "meta 1",
				},
			},
			want: "Error #01: error 1\n     Meta: meta 1\n",
		},
		{
			name: "multiple",
			a: errorMsgs{
				{
					Err:  errors.New("error 1"),
					Meta: "meta 1",
				},
				{
					Err:  errors.New("error 2"),
					Meta: "meta 2",
				},
			},
			want: "Error #01: error 1\n     Meta: meta 1\nError #02: error 2\n     Meta: meta 2\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.a.String(); got != tt.want {
				t.Errorf("errorMsgs.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

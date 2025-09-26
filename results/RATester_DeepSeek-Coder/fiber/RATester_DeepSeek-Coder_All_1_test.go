package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAll_1(t *testing.T) {
	app := New()
	registering := &Registering{app: app, path: "/test"}

	handler := func(c Ctx) error {
		return c.SendString("Hello, World 👋")
	}

	tests := []struct {
		name string
		r    *Registering
		want *Registering
	}{
		{
			name: "TestAll",
			r:    registering,
			want: registering,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.r.All(handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

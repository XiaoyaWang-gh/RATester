package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func Teststack_1(t *testing.T) {
	tests := []struct {
		name string
		skip int
		want []byte
	}{
		{
			name: "Test case 1",
			skip: 0,
			want: []byte("test.go:123 (0x456)\n\tfunction(): source line\n"),
		},
		{
			name: "Test case 2",
			skip: 1,
			want: []byte("test.go:456 (0x789)\n\tfunction(): source line\n"),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := stack(tt.skip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack() = %v, want %v", got, tt.want)
			}
		})
	}
}

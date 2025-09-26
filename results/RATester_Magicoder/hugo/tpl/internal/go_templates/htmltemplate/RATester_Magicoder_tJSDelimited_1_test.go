package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttJSDelimited_1(t *testing.T) {
	tests := []struct {
		name   string
		input  context
		input2 []byte
		want   context
		want2  int
	}{
		{
			name: "Test case 1",
			input: context{
				state: stateJSSqStr,
			},
			input2: []byte(`"Hello, world!"`),
			want: context{
				state: stateJSSqStr,
			},
			want2: 15,
		},
		{
			name: "Test case 2",
			input: context{
				state: stateJSRegexp,
			},
			input2: []byte(`/Hello, world!/`),
			want: context{
				state: stateJSRegexp,
			},
			want2: 16,
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

			got, got2 := tJSDelimited(tt.input, tt.input2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tJSDelimited() got = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("tJSDelimited() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

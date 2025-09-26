package hints

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHint_1(t *testing.T) {
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want *Hint
	}{
		{
			name: "TestNewHint_0",
			args: args{
				key:   "test_key",
				value: "test_value",
			},
			want: &Hint{
				key:   "test_key",
				value: "test_value",
			},
		},
		{
			name: "TestNewHint_1",
			args: args{
				key:   123,
				value: 456,
			},
			want: &Hint{
				key:   123,
				value: 456,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewHint(tt.args.key, tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHint() = %v, want %v", got, tt.want)
			}
		})
	}
}

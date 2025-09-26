package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValue_3(t *testing.T) {
	type args struct {
		source []byte
		lh     LowHigh[[]byte]
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test case 1",
			args: args{
				source: []byte("hello world"),
				lh:     LowHigh[[]byte]{Low: 0, High: 5},
			},
			want: []byte("hello"),
		},
		{
			name: "Test case 2",
			args: args{
				source: []byte("hello world"),
				lh:     LowHigh[[]byte]{Low: 6, High: 11},
			},
			want: []byte("world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.lh.Value(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

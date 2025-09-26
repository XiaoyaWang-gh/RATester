package hints

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRelDepth_1(t *testing.T) {
	type args struct {
		d int
	}
	tests := []struct {
		name string
		args args
		want *Hint
	}{
		{
			name: "Test case 1",
			args: args{d: 1},
			want: &Hint{KeyRelDepth, 1},
		},
		{
			name: "Test case 2",
			args: args{d: 2},
			want: &Hint{KeyRelDepth, 2},
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

			if got := RelDepth(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RelDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

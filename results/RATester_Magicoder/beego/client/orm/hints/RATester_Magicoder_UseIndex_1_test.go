package hints

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUseIndex_1(t *testing.T) {
	type args struct {
		indexes []string
	}
	tests := []struct {
		name string
		args args
		want *Hint
	}{
		{
			name: "TestUseIndex_0",
			args: args{
				indexes: []string{"index1", "index2"},
			},
			want: &Hint{
				key:   "UseIndex",
				value: []string{"index1", "index2"},
			},
		},
		{
			name: "TestUseIndex_1",
			args: args{
				indexes: []string{"index3", "index4"},
			},
			want: &Hint{
				key:   "UseIndex",
				value: []string{"index3", "index4"},
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

			got := UseIndex(tt.args.indexes...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

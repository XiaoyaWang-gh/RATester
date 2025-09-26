package hints

import (
	"fmt"
	"reflect"
	"testing"
)

func TestForceIndex_1(t *testing.T) {
	type args struct {
		indexes []string
	}
	tests := []struct {
		name string
		args args
		want *Hint
	}{
		{
			name: "TestForceIndex_0",
			args: args{
				indexes: []string{"index1", "index2"},
			},
			want: &Hint{
				key:   KeyForceIndex,
				value: []string{"index1", "index2"},
			},
		},
		{
			name: "TestForceIndex_1",
			args: args{
				indexes: []string{"index3", "index4"},
			},
			want: &Hint{
				key:   KeyForceIndex,
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

			if got := ForceIndex(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForceIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetData_1(t *testing.T) {
	type args struct {
		key interface{}
		val interface{}
	}
	tests := []struct {
		name  string
		input *BeegoInput
		args  args
		want  map[interface{}]interface{}
	}{
		{
			name: "TestSetData",
			input: &BeegoInput{
				data: make(map[interface{}]interface{}),
			},
			args: args{
				key: "key1",
				val: "value1",
			},
			want: map[interface{}]interface{}{
				"key1": "value1",
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

			tt.input.SetData(tt.args.key, tt.args.val)
			if !reflect.DeepEqual(tt.input.data, tt.want) {
				t.Errorf("got %v, want %v", tt.input.data, tt.want)
			}
		})
	}
}

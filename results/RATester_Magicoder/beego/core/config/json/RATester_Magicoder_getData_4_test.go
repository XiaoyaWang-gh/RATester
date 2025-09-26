package json

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetData_4(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Test case 1",
			args: args{
				key: "section::key",
			},
			want: "value",
		},
		{
			name: "Test case 2",
			args: args{
				key: "section::key2",
			},
			want: "value2",
		},
		{
			name: "Test case 3",
			args: args{
				key: "section::key3",
			},
			want: "value3",
		},
		{
			name: "Test case 4",
			args: args{
				key: "section::key4",
			},
			want: "value4",
		},
		{
			name: "Test case 5",
			args: args{
				key: "section::key5",
			},
			want: "value5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &JSONConfigContainer{
				data: map[string]interface{}{
					"section": map[string]interface{}{
						"key":  "value",
						"key2": "value2",
						"key3": "value3",
						"key4": "value4",
						"key5": "value5",
					},
				},
			}
			if got := c.getData(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getData() = %v, want %v", got, tt.want)
			}
		})
	}
}

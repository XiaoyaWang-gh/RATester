package binding

import (
	"fmt"
	"testing"
)

func TestsetFormMap_1(t *testing.T) {
	type args struct {
		ptr  any
		form map[string][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				ptr: &map[string]string{},
				form: map[string][]string{
					"key1": {"value1"},
					"key2": {"value2"},
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				ptr: &map[string][]string{},
				form: map[string][]string{
					"key1": {"value1", "value2"},
					"key2": {"value3", "value4"},
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				ptr: &map[string]string{},
				form: map[string][]string{
					"key1": {"value1", "value2"},
					"key2": {"value3", "value4"},
				},
			},
			wantErr: true,
		},
		{
			name: "Test case 4",
			args: args{
				ptr: &map[string]string{},
				form: map[string][]string{
					"key1": {"value1"},
					"key2": {"value2"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := setFormMap(tt.args.ptr, tt.args.form); (err != nil) != tt.wantErr {
				t.Errorf("setFormMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

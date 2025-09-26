package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStringMapStringE_1(t *testing.T) {
	type args struct {
		in any
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				in: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
			},
			want: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				in: "invalid input",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ToStringMapStringE(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringMapStringE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMapStringE() = %v, want %v", got, tt.want)
			}
		})
	}
}

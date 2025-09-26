package middleware

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValuesFromForm_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				name: "test_name",
			},
			want:    []string{"test_value"},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				name: "test_name_2",
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

			got := valuesFromForm(tt.args.name)
			got1, err := got(nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("valuesFromForm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got1, tt.want) {
				t.Errorf("valuesFromForm() got1 = %v, want %v", got1, tt.want)
			}
		})
	}
}

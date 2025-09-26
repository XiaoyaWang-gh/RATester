package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func Testtrim_1(t *testing.T) {
	type args struct {
		name string
		key  string
		s    []string
	}
	tests := []struct {
		name    string
		args    args
		wantTs  []interface{}
		wantErr bool
	}{
		{
			name: "Test trim function",
			args: args{
				name: "test",
				key:  "key",
				s:    []string{"value1", "value2"},
			},
			wantTs:  []interface{}{"value1", "value2", "key"},
			wantErr: false,
		},
		{
			name: "Test trim function with error",
			args: args{
				name: "test",
				key:  "key",
				s:    []string{"value1", "value2"},
			},
			wantTs:  []interface{}{"value1", "value2", "key"},
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

			gotTs, err := trim(tt.args.name, tt.args.key, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("trim() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTs, tt.wantTs) {
				t.Errorf("trim() = %v, want %v", gotTs, tt.wantTs)
			}
		})
	}
}

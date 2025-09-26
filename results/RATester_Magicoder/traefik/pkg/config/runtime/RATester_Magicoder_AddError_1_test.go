package runtime

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestAddError_1(t *testing.T) {
	type args struct {
		err      error
		critical bool
	}
	tests := []struct {
		name string
		s    *UDPServiceInfo
		args args
		want *UDPServiceInfo
	}{
		{
			name: "add error",
			s:    &UDPServiceInfo{Err: []string{"error1"}},
			args: args{
				err:      errors.New("error2"),
				critical: false,
			},
			want: &UDPServiceInfo{Err: []string{"error1", "error2"}},
		},
		{
			name: "add critical error",
			s:    &UDPServiceInfo{Err: []string{"error1"}},
			args: args{
				err:      errors.New("error2"),
				critical: true,
			},
			want: &UDPServiceInfo{Err: []string{"error1", "error2"}, Status: StatusDisabled},
		},
		{
			name: "add existing error",
			s:    &UDPServiceInfo{Err: []string{"error1"}},
			args: args{
				err:      errors.New("error1"),
				critical: false,
			},
			want: &UDPServiceInfo{Err: []string{"error1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.s.AddError(tt.args.err, tt.args.critical)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("got %v, want %v", tt.s, tt.want)
			}
		})
	}
}

package try

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHasHeader_1(t *testing.T) {
	type args struct {
		header string
	}
	tests := []struct {
		name string
		args args
		want ResponseCondition
	}{
		{
			name: "Test case 1",
			args: args{
				header: "Content-Type",
			},
			want: func(res *http.Response) error {
				if _, ok := res.Header["Content-Type"]; !ok {
					return errors.New("response doesn't contain header: Content-Type")
				}
				return nil
			},
		},
		{
			name: "Test case 2",
			args: args{
				header: "Authorization",
			},
			want: func(res *http.Response) error {
				if _, ok := res.Header["Authorization"]; !ok {
					return errors.New("response doesn't contain header: Authorization")
				}
				return nil
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

			if got := HasHeader(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

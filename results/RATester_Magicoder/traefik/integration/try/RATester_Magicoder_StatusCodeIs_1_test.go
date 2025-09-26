package try

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStatusCodeIs_1(t *testing.T) {
	type args struct {
		status int
	}
	tests := []struct {
		name string
		args args
		want ResponseCondition
	}{
		{
			name: "TestStatusCodeIs_200",
			args: args{status: 200},
			want: func(res *http.Response) error {
				if res.StatusCode != 200 {
					return fmt.Errorf("got status code %d, wanted %d", res.StatusCode, 200)
				}
				return nil
			},
		},
		{
			name: "TestStatusCodeIs_404",
			args: args{status: 404},
			want: func(res *http.Response) error {
				if res.StatusCode != 404 {
					return fmt.Errorf("got status code %d, wanted %d", res.StatusCode, 404)
				}
				return nil
			},
		},
		{
			name: "TestStatusCodeIs_500",
			args: args{status: 500},
			want: func(res *http.Response) error {
				if res.StatusCode != 500 {
					return fmt.Errorf("got status code %d, wanted %d", res.StatusCode, 500)
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

			if got := StatusCodeIs(tt.args.status); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusCodeIs() = %v, want %v", got, tt.want)
			}
		})
	}
}

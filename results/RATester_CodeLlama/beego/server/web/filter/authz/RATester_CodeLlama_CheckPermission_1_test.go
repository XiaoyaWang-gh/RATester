package authz

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCheckPermission_1(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		a    *BasicAuthorizer
		args args
		want bool
	}{
		{
			name: "case 1",
			a:    &BasicAuthorizer{},
			args: args{r: &http.Request{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.a.CheckPermission(tt.args.r); got != tt.want {
				t.Errorf("BasicAuthorizer.CheckPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

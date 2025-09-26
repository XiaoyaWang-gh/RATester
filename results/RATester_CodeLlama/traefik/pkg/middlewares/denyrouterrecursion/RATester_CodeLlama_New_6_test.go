package denyrouterrecursion

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNew_6(t *testing.T) {
	type args struct {
		routerName string
		next       http.Handler
	}
	tests := []struct {
		name    string
		args    args
		want    *DenyRouterRecursion
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				routerName: "test",
				next:       nil,
			},
			want: &DenyRouterRecursion{
				routerName:     "test",
				routerNameHash: "test",
				next:           nil,
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				routerName: "",
				next:       nil,
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

			got, err := New(tt.args.routerName, tt.args.next)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

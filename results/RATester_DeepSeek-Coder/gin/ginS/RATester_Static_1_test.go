package ginS

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStatic_1(t *testing.T) {
	type args struct {
		relativePath string
		root         string
	}
	tests := []struct {
		name string
		args args
		want gin.IRoutes
	}{
		{
			name: "Test Case 1",
			args: args{
				relativePath: "relativePath",
				root:         "root",
			},
			want: Static("relativePath", "root"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Static(tt.args.relativePath, tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Static() = %v, want %v", got, tt.want)
			}
		})
	}
}

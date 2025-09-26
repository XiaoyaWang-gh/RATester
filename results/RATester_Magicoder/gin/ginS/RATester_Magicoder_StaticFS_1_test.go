package ginS

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStaticFS_1(t *testing.T) {
	type args struct {
		relativePath string
		fs           http.FileSystem
	}
	tests := []struct {
		name string
		args args
		want gin.IRoutes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := StaticFS(tt.args.relativePath, tt.args.fs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StaticFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

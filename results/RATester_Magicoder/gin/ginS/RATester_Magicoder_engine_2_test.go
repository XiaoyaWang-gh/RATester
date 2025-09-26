package ginS

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func Testengine_2(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		{
			name: "Test engine",
			want: engine(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := engine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("engine() = %v, want %v", got, tt.want)
			}
		})
	}
}

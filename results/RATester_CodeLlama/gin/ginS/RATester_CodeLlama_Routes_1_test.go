package ginS

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRoutes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var expected gin.RoutesInfo
	var actual gin.RoutesInfo

	// Act
	actual = Routes()

	// Assert
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

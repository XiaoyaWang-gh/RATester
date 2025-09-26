package ginS

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticFS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	relativePath := "test"
	fs := http.Dir("test")

	// Act
	engine := StaticFS(relativePath, fs)

	// Assert
	assert.NotNil(t, engine)
}

package observability

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestNewMiddleware_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var next http.Handler
	var name string
	var typeName string
	var spanKind trace.SpanKind

	// Act
	var actual = NewMiddleware(next, name, typeName, spanKind)

	// Assert
	assert.NotNil(t, actual)
}

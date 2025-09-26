package recover //nolint:predeclared // TODO: Rename to some non-builtin
import (
	"fmt"
	"testing"
)

func TestDefaultStackTraceHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}

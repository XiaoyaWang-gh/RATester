package markup

import (
	"fmt"
	"testing"
)

func TestNewConverterProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// TODO(bep)
}

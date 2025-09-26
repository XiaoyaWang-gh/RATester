package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestNewOxyWrapper_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	ow := NewOxyWrapper(logger)

	if ow == nil {
		t.Errorf("NewOxyWrapper() = %v, want %v", ow, "not nil")
	}
}

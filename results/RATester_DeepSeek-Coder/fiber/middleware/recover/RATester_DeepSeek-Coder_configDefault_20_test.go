package recover //nolint:predeclared // TODO: Rename to some non-builtin
import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
	"gotest.tools/assert"
)

func TestConfigDefault_20(t *testing.T) {
	t.Parallel()

	defaultHandler := func(c fiber.Ctx, e any) {}

	tests := []struct {
		name     string
		config   []Config
		expected Config
	}{
		{
			name: "no config",
			expected: Config{
				EnableStackTrace:  false,
				StackTraceHandler: defaultStackTraceHandler,
			},
		},
		{
			name: "with config",
			config: []Config{
				{
					EnableStackTrace:  true,
					StackTraceHandler: defaultHandler,
				},
			},
			expected: Config{
				EnableStackTrace:  true,
				StackTraceHandler: defaultHandler,
			},
		},
		{
			name: "with config and no stack trace handler",
			config: []Config{
				{
					EnableStackTrace: true,
				},
			},
			expected: Config{
				EnableStackTrace:  true,
				StackTraceHandler: defaultStackTraceHandler,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := configDefault(test.config...)
			assert.Equal(t, test.expected, result)
		})
	}
}

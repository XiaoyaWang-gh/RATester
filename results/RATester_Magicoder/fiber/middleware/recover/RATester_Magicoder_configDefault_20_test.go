package recover //nolint:predeclared // TODO: Rename to some non-builtin
import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestconfigDefault_20(t *testing.T) {
	tests := []struct {
		name string
		args []Config
		want Config
	}{
		{
			name: "Test with no args",
			args: []Config{},
			want: ConfigDefault,
		},
		{
			name: "Test with args",
			args: []Config{
				{
					EnableStackTrace: true,
					StackTraceHandler: func(c fiber.Ctx, e any) {
						// handle stack trace
					},
				},
			},
			want: Config{
				EnableStackTrace: true,
				StackTraceHandler: func(c fiber.Ctx, e any) {
					// handle stack trace
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := configDefault(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

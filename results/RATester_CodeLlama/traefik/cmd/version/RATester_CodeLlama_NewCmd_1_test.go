package version

import (
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/paerser/cli"
)

func TestNewCmd_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		expected *cli.Command
	}{
		{
			name: "NewCmd",
			expected: &cli.Command{
				Name:          "version",
				Description:   "Shows the current Traefik version.",
				Configuration: nil,
				Run: func(_ []string) error {
					if err := GetPrint(os.Stdout); err != nil {
						return err
					}
					fmt.Print("\n")
					return nil
				},
			},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			result := NewCmd()

			assert.Equal(t, test.expected, result)
		})
	}
}

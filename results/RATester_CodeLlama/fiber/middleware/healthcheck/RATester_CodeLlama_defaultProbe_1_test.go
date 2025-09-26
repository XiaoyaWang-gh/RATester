package healthcheck

import (
	"fmt"
	"testing"
)

func TestDefaultProbe_1(t *testing.T) {
	// Uncomment this line to follow test logs
	// log.SetOutput(os.Stdout)

	// Your test cases
	t.Run("Testcase 1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// Setup
		// Test variables
		// Run your test
		// Assert results
	})
}

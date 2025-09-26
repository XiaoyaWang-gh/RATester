package sub

import (
	"fmt"
	"testing"
)

func TestRunClient_1(t *testing.T) {
	t.Run("test runClient", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := runClient("")
		if err != nil {
			t.Errorf("runClient() error = %v", err)
		}
	})
}

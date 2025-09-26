package logs

import (
	"fmt"
	"testing"
)

func TestDestroy_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	writer := &SMTPWriter{
		Username: "testuser",
		Password: "testpassword",
		Host:     "testhost",
	}
	writer.Destroy()
}

package logs

import (
	"fmt"
	"testing"
)

func TestDestroy_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &multiFileLogWriter{}
	f.Destroy()
}

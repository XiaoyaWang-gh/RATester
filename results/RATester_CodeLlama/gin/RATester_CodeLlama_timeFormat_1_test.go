package gin

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFormat_1(t *testing.T) {
	t.Run("test timeFormat", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t1 := time.Now()
		want := t1.Format("2006/01/02 - 15:04:05")
		got := timeFormat(t1)
		if got != want {
			t.Errorf("timeFormat() = %v, want %v", got, want)
		}
	})
}

package mock

import (
	"fmt"
	"testing"
)

func TestGroupBy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d *DoNothingQuerySetter
	exprs := []string{"exprs"}
	d.GroupBy(exprs...)
}

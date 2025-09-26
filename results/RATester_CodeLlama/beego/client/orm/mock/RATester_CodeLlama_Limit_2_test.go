package mock

import (
	"fmt"
	"testing"
)

func TestLimit_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d *DoNothingQuerySetter
	var limit interface{}
	var args []interface{}
	d.Limit(limit, args...)
}

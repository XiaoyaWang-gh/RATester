package metric

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDec_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	counter := &StandardDateCounter{
		reserveDays:    7,
		counts:         make([]int64, 7),
		lastUpdateDate: time.Now(),
	}

	counter.Inc(10)
	counter.Dec(5)

	expected := []int64{0, 0, 0, 0, 0, 0, 0}
	actual := counter.GetLastDaysCount(7)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

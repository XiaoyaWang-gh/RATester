package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/logg"
)

func TestLoggFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &buildCounters{}
	c.pageRenderCounter.Store(1)
	c.contentRenderCounter.Store(2)
	want := logg.Fields{
		{Name: "pages", Value: 1},
		{Name: "content", Value: 2},
	}
	if got := c.loggFields(); !reflect.DeepEqual(got, want) {
		t.Errorf("loggFields() = %v, want %v", got, want)
	}
}

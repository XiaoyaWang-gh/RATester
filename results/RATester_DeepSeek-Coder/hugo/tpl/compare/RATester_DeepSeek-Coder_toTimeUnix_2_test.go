package compare

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestToTimeUnix_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{loc: time.UTC}
	v := reflect.ValueOf(time.Now())
	unixTime := ns.toTimeUnix(v)
	if unixTime != v.Interface().(time.Time).Unix() {
		t.Errorf("Expected %v, got %v", v.Interface().(time.Time).Unix(), unixTime)
	}
}

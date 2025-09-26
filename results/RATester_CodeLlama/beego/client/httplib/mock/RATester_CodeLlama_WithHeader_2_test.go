package mock

import (
	"fmt"
	"net/textproto"
	"testing"
)

func TestWithHeader_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "key"
	value := "value"
	sc := &SimpleCondition{}
	WithHeader(key, value)(sc)
	if sc.header[textproto.CanonicalMIMEHeaderKey(key)] != value {
		t.Errorf("WithHeader failed")
	}
}

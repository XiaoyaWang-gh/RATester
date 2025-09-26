package logs

import (
	"fmt"
	"testing"
)

func TestDebugf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := ElasticLogger{}
	l.Debugf("test")
}

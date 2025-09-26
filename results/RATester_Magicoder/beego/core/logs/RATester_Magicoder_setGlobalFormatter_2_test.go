package logs

import (
	"fmt"
	"testing"
)

func TestsetGlobalFormatter_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{}
	testFormatter := "testFormatter"
	err := bl.setGlobalFormatter(testFormatter)
	if err != nil {
		t.Errorf("setGlobalFormatter failed, error: %v", err)
	}
	if bl.globalFormatter != testFormatter {
		t.Errorf("setGlobalFormatter failed, expected: %s, got: %s", testFormatter, bl.globalFormatter)
	}
}

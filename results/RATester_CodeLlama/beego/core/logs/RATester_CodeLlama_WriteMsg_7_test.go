package logs

import (
	"fmt"
	"testing"
)

func TestWriteMsg_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SLACKWriter{}
	lm := &LogMsg{}
	err := s.WriteMsg(lm)
	if err != nil {
		t.Errorf("TestWriteMsg error %v", err)
	}
}

package framework

import (
	"fmt"
	"testing"
)

func TestNowStamp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nowStamp()
	if nowStamp() != "2019/01/01 12:00:00.000" {
		t.Errorf("nowStamp() = %v, want %v", nowStamp(), "2019/01/01 12:00:00.000")
	}
}

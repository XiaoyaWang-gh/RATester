package orm

import (
	"fmt"
	"testing"
)

func TestDelete_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	orm := &DoNothingOrm{}
	_, err := orm.Delete(nil)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
}

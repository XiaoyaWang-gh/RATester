package page

import (
	"fmt"
	"testing"
)

func TestGetPage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	_, err := nop.GetPage("")
	if err != nil {
		t.Errorf("GetPage() error = %v, wantErr %v", err, nil)
		return
	}
}

package page

import (
	"fmt"
	"testing"
)

func TestPaginate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	_, err := nop.Paginate(nil)
	if err != nil {
		t.Errorf("Paginate() error = %v, wantErr %v", err, nil)
		return
	}
}

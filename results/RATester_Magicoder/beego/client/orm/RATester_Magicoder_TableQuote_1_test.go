package orm

import (
	"fmt"
	"testing"
)

func TestTableQuote_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBase{}
	expected := "`"
	actual := db.TableQuote()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

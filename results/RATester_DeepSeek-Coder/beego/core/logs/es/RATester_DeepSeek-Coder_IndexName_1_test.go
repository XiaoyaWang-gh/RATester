package es

import (
	"fmt"
	"testing"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func TestIndexName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &defaultIndexNaming{}
	lm := &logs.LogMsg{
		When: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	expected := "2022.01.01"
	result := d.IndexName(lm)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

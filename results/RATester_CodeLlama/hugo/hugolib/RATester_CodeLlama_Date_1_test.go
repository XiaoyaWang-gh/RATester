package hugolib

import (
	"fmt"
	"testing"
	"time"

	"github.com/gohugoio/hugo/resources/page/pagemeta"
)

func TestDate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig = &pagemeta.PageConfig{}
	p.pageConfig.Dates.Date = time.Now()
	if p.Date().IsZero() {
		t.Errorf("Date is zero")
	}
}

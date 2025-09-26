package pagemeta

import (
	"fmt"
	"testing"
)

func TestUpdateDateAndLastmodAndPublishDateIfAfter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Dates{}
	in := Dates{}
	d.UpdateDateAndLastmodAndPublishDateIfAfter(in)
	if d.Date != in.Date {
		t.Errorf("d.Date = %v, want %v", d.Date, in.Date)
	}
	if d.Lastmod != in.Lastmod {
		t.Errorf("d.Lastmod = %v, want %v", d.Lastmod, in.Lastmod)
	}
	if d.PublishDate != in.PublishDate {
		t.Errorf("d.PublishDate = %v, want %v", d.PublishDate, in.PublishDate)
	}
}

package exif

import (
	"fmt"
	"regexp"
	"testing"
)

func TestShouldExclude_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Decoder{}
	d.excludeFieldsrRe = regexp.MustCompile("^foo$")
	if d.shouldExclude("foo") != true {
		t.Errorf("shouldExclude failed")
	}
	if d.shouldExclude("bar") != false {
		t.Errorf("shouldExclude failed")
	}
}

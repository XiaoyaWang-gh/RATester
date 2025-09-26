package pagemeta

import (
	"fmt"
	"testing"
	"time"
)

func TestcreateDateHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fm := FrontMatterHandler{}
	identifiers := []string{fmFilename, ":modTime", ":gitAuthorDate"}
	setter := func(d *FrontMatterDescriptor, t time.Time) {}
	_, err := fm.createDateHandler(identifiers, setter)
	if err != nil {
		t.Errorf("createDateHandler() error = %v", err)
		return
	}
}

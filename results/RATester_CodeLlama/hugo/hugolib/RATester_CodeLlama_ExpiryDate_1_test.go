package hugolib

import (
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestExpiryDate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig.Dates.ExpiryDate = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, p.ExpiryDate(), time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
}

package orm

import (
	"fmt"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
	"github.com/zeebo/assert"
)

func TestCountSQL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBase{}
	qs := querySet{}
	mi := &models.ModelInfo{}
	cond := &Condition{}
	tz := &time.Location{}
	got, _ := d.countSQL(qs, mi, cond, tz)
	want := ""
	assert.Equal(t, want, got)
}

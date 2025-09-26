package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestReadOrCreate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	md := &models.ModelInfo{}
	col1 := "col1"
	cols := []string{"col1", "col2"}
	b, id, err := o.ReadOrCreate(md, col1, cols...)
	if err != nil {
		t.Errorf("ReadOrCreate error: %v", err)
	}
	if b {
		t.Errorf("ReadOrCreate should return false")
	}
	if id != 0 {
		t.Errorf("ReadOrCreate should return 0")
	}
}

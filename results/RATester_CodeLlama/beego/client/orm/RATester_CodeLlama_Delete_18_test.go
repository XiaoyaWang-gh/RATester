package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestDelete_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	md := &models.ModelInfo{}
	cols := []string{"col1", "col2"}
	_, err := o.Delete(md, cols...)
	if err != nil {
		t.Errorf("Delete() error = %v", err)
		return
	}
}

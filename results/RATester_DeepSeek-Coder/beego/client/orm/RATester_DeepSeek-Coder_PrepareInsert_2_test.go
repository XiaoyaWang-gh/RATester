package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestPrepareInsert_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{
		mi: &models.ModelInfo{},
	}

	inserter, err := o.PrepareInsert()
	if err != nil {
		t.Errorf("PrepareInsert() error = %v", err)
		return
	}

	if inserter == nil {
		t.Errorf("PrepareInsert() = %v, want not nil", inserter)
	}
}

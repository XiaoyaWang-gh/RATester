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
		t.Errorf("Error preparing insert: %v", err)
	}

	_, err = inserter.Insert(struct{}{})
	if err != nil {
		t.Errorf("Error inserting: %v", err)
	}

	err = inserter.Close()
	if err != nil {
		t.Errorf("Error closing inserter: %v", err)
	}
}

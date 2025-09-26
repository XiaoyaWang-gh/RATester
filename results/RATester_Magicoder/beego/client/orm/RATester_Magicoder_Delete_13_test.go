package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestDelete_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{
		mi: &models.ModelInfo{},
	}
	count, err := o.Delete()
	if err != nil {
		t.Errorf("Delete() error = %v", err)
		return
	}
	if count != 1 {
		t.Errorf("Delete() count = %v, want %v", count, 1)
	}
}

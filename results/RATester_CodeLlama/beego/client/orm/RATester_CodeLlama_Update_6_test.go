package orm

import (
	"fmt"
	"testing"

	"github.com/minio/console/models"
)

func TestUpdate_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	md := &models.User{}
	cols := []string{"name"}
	_, err := o.Update(md, cols...)
	if err != nil {
		t.Errorf("TestUpdate error, err:%v", err)
	}
}

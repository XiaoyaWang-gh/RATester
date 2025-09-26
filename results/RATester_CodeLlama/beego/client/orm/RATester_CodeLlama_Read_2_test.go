package orm

import (
	"fmt"
	"testing"

	"github.com/minio/console/models"
)

func TestRead_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	md := &models.User{}
	cols := []string{"id", "name"}
	err := o.Read(md, cols...)
	if err != nil {
		t.Errorf("Read() error = %v", err)
		return
	}
}

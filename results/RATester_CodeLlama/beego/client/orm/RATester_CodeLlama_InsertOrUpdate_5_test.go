package orm

import (
	"fmt"
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/model"
)

func TestInsertOrUpdate_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingOrm{}
	md := &model.User{}
	colConflitAndArgs := []string{"id"}
	_, err := d.InsertOrUpdate(md, colConflitAndArgs...)
	if err != nil {
		t.Errorf("InsertOrUpdate() error = %v", err)
		return
	}
}

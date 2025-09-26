package migration

import (
	"fmt"
	"testing"
)

func TestGetSQL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{
		sqls:           []string{},
		Created:        "",
		TableName:      "",
		Engine:         "",
		Charset:        "",
		ModifyType:     "",
		Columns:        []*Column{},
		Indexes:        []*Index{},
		Primary:        []*Column{},
		Uniques:        []*Unique{},
		Foreigns:       []*Foreign{},
		Renames:        []*RenameColumn{},
		RemoveColumns:  []*Column{},
		RemoveIndexes:  []*Index{},
		RemoveUniques:  []*Unique{},
		RemoveForeigns: []*Foreign{},
	}
	sql := m.GetSQL()
	if sql != "" {
		t.Errorf("GetSQL() = %v, want %v", sql, "")
	}
}

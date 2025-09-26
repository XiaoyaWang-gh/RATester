package orm

import (
	"fmt"
	"testing"
)

func TestGenerateSpecifyIndex_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d dbBaseSqlite
	var tableName string
	var useIndex int
	var indexes []string
	var want string
	var got string
	tableName = "tableName"
	useIndex = 1
	indexes = []string{"index1", "index2"}
	want = " INDEXED BY index1,index2 "
	got = d.GenerateSpecifyIndex(tableName, useIndex, indexes)
	if got != want {
		t.Errorf("GenerateSpecifyIndex() = %v, want %v", got, want)
	}
}

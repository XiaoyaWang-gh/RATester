package orm

import (
	"fmt"
	"testing"
)

func TestGenerateSpecifyIndex_3(t *testing.T) {
	type args struct {
		tableName string
		useIndex  int
		indexes   []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &dbBasePostgres{}
			if got := d.GenerateSpecifyIndex(tt.args.tableName, tt.args.useIndex, tt.args.indexes); got != tt.want {
				t.Errorf("dbBasePostgres.GenerateSpecifyIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

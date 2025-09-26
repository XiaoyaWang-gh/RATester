package codegen

import (
	"fmt"
	"testing"
)

func TestTypeName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "github.com/go-sql-driver/mysql.MySQLDriver"
	pkg := "github.com/go-sql-driver/mysql"
	want := "MySQLDriver"
	got := typeName(name, pkg)
	if got != want {
		t.Errorf("typeName(%q, %q) = %q; want %q", name, pkg, got, want)
	}
}

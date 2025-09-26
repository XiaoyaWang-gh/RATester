package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestUpdate_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type User struct {
		Id   int64
		Name string
	}

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	o := &ormBase{db: db}

	_, err = o.Insert(&User{Name: "John"})
	if err != nil {
		t.Fatal(err)
	}

	_, err = o.Update(&User{Id: 1, Name: "Jane"})
	if err != nil {
		t.Fatal(err)
	}

	var user User
	err = o.Read(&user)
	if err != nil {
		t.Fatal(err)
	}

	if user.Name != "Jane" {
		t.Errorf("Expected Name to be 'Jane', got '%s'", user.Name)
	}
}

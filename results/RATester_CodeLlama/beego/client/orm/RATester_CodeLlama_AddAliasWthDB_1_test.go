package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestAddAliasWthDB_1(t *testing.T) {
	t.Run("test case 1:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// arrange
		var aliasName string
		var driverName string
		var db *sql.DB
		var params []DBOption
		// act
		err := AddAliasWthDB(aliasName, driverName, db, params...)
		// assert
		if err != nil {
			t.Errorf("err is not nil, err: %v", err)
			return
		}
	})
}

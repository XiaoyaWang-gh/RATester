package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestGetDB_1(t *testing.T) {
	tests := []struct {
		name        string
		aliasNames  []string
		expectedDB  *sql.DB
		expectedErr error
	}{
		{
			name:        "Test with alias name",
			aliasNames:  []string{"test"},
			expectedDB:  &sql.DB{},
			expectedErr: nil,
		},
		{
			name:        "Test with no alias name",
			aliasNames:  []string{},
			expectedDB:  &sql.DB{},
			expectedErr: nil,
		},
		{
			name:        "Test with non-existent alias name",
			aliasNames:  []string{"non-existent"},
			expectedDB:  nil,
			expectedErr: fmt.Errorf("DataBase of alias name `non-existent` not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			db, err := GetDB(test.aliasNames...)
			if err != nil && err.Error() != test.expectedErr.Error() {
				t.Errorf("Expected error %v, but got %v", test.expectedErr, err)
			}
			if db != test.expectedDB {
				t.Errorf("Expected DB %v, but got %v", test.expectedDB, db)
			}
		})
	}
}

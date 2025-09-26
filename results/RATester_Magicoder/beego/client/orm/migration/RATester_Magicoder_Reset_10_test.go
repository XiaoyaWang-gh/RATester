package migration

import (
	"fmt"
	"testing"
)

func TestReset_10(t *testing.T) {
	type testCase struct {
		name    string
		m       Migration
		wantErr bool
	}

	testCases := []testCase{
		{
			name: "test case 1",
			m: Migration{
				sqls: []string{
					"CREATE TABLE test_table (id INT, name VARCHAR(255))",
				},
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			m: Migration{
				sqls: []string{
					"CREATE TABLE test_table (id INT, name VARCHAR(255))",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := Reset()
			if (err != nil) != tc.wantErr {
				t.Errorf("Reset() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

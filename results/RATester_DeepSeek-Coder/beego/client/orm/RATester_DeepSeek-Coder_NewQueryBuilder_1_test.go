package orm

import (
	"fmt"
	"testing"
)

func TestNewQueryBuilder_1(t *testing.T) {
	tests := []struct {
		name    string
		driver  string
		wantErr bool
	}{
		{"Test MySQL driver", "mysql", false},
		{"Test TiDB driver", "tidb", false},
		{"Test Postgres driver", "postgres", false},
		{"Test SQLite driver", "sqlite", true},
		{"Test unknown driver", "unknown", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := NewQueryBuilder(tt.driver)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQueryBuilder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
		{
			name:    "Test MySQL",
			driver:  "mysql",
			wantErr: false,
		},
		{
			name:    "Test TiDB",
			driver:  "tidb",
			wantErr: false,
		},
		{
			name:    "Test Postgres",
			driver:  "postgres",
			wantErr: false,
		},
		{
			name:    "Test SQLite",
			driver:  "sqlite",
			wantErr: true,
		},
		{
			name:    "Test Unknown Driver",
			driver:  "unknown",
			wantErr: true,
		},
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
				return
			}
		})
	}
}

package orm

import (
	"fmt"
	"testing"
)

func TestRun_4(t *testing.T) {
	tests := []struct {
		name    string
		d       *commandSQLAll
		wantErr bool
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

			d := &commandSQLAll{
				al: tt.d.al,
			}
			if err := d.Run(); (err != nil) != tt.wantErr {
				t.Errorf("commandSQLAll.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

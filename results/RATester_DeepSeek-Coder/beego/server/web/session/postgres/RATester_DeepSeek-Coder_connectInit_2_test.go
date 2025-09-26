package postgres

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestConnectInit_2(t *testing.T) {
	type fields struct {
		maxlifetime int64
		savePath    string
	}
	tests := []struct {
		name   string
		fields fields
		want   *sql.DB
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

			mp := &Provider{
				maxlifetime: tt.fields.maxlifetime,
				savePath:    tt.fields.savePath,
			}
			if got := mp.connectInit(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.connectInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

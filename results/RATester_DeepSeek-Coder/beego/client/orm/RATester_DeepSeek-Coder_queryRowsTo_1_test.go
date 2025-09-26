package orm

import (
	"fmt"
	"testing"
)

func TestQueryRowsTo_1(t *testing.T) {
	type args struct {
		container interface{}
		keyCol    string
		valueCol  string
	}
	tests := []struct {
		name    string
		o       *rawSet
		args    args
		want    int64
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

			got, err := tt.o.queryRowsTo(tt.args.container, tt.args.keyCol, tt.args.valueCol)
			if (err != nil) != tt.wantErr {
				t.Errorf("rawSet.queryRowsTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("rawSet.queryRowsTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

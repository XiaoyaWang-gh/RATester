package orm

import (
	"fmt"
	"testing"
)

func TestreadValues_2(t *testing.T) {
	type args struct {
		container interface{}
		needCols  []string
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

			got, err := tt.o.readValues(tt.args.container, tt.args.needCols)
			if (err != nil) != tt.wantErr {
				t.Errorf("rawSet.readValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("rawSet.readValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd_9(t *testing.T) {
	type args struct {
		owner        *moduleAdapter
		moduleImport Import
	}
	tests := []struct {
		name    string
		c       *collector
		args    args
		want    *moduleAdapter
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

			got, err := tt.c.add(tt.args.owner, tt.args.moduleImport)
			if (err != nil) != tt.wantErr {
				t.Errorf("collector.add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collector.add() = %v, want %v", got, tt.want)
			}
		})
	}
}

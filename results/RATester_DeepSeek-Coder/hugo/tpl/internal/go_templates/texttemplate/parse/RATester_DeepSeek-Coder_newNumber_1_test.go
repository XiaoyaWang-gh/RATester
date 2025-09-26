package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewNumber_1(t *testing.T) {
	type args struct {
		pos  Pos
		text string
		typ  itemType
	}
	tests := []struct {
		name    string
		args    args
		want    *NumberNode
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

			tr := &Tree{}
			got, err := tr.newNumber(tt.args.pos, tt.args.text, tt.args.typ)
			if (err != nil) != tt.wantErr {
				t.Errorf("newNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

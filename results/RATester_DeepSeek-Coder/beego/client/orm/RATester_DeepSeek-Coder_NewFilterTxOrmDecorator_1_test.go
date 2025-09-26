package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFilterTxOrmDecorator_1(t *testing.T) {
	type args struct {
		delegate TxOrmer
		root     Filter
		txName   string
	}
	tests := []struct {
		name string
		args args
		want TxOrmer
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

			if got := NewFilterTxOrmDecorator(tt.args.delegate, tt.args.root, tt.args.txName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterTxOrmDecorator() = %v, want %v", got, tt.want)
			}
		})
	}
}

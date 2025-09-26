package orm

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestBegin_1(t *testing.T) {
	type fields struct {
		ormer       ormer
		TxBeginner  TxBeginner
		TxCommitter TxCommitter
		root        Filter
		insideTx    bool
		txStartTime time.Time
		txName      string
	}
	tests := []struct {
		name    string
		fields  fields
		want    TxOrmer
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

			f := &filterOrmDecorator{
				ormer:       tt.fields.ormer,
				TxBeginner:  tt.fields.TxBeginner,
				TxCommitter: tt.fields.TxCommitter,
				root:        tt.fields.root,
				insideTx:    tt.fields.insideTx,
				txStartTime: tt.fields.txStartTime,
				txName:      tt.fields.txName,
			}
			got, err := f.Begin()
			if (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.Begin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterOrmDecorator.Begin() = %v, want %v", got, tt.want)
			}
		})
	}
}

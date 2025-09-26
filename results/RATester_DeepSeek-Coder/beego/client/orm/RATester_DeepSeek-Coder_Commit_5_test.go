package orm

import (
	"fmt"
	"testing"
	"time"
)

func TestCommit_5(t *testing.T) {
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
			if err := f.Commit(); (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.Commit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

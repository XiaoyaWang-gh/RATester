package logs

import (
	"fmt"
	"testing"
)

func TestSetGlobalFormatter_2(t *testing.T) {
	type args struct {
		fmtter string
	}
	tests := []struct {
		name    string
		bl      *BeeLogger
		args    args
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

			bl := &BeeLogger{}
			if err := bl.setGlobalFormatter(tt.args.fmtter); (err != nil) != tt.wantErr {
				t.Errorf("BeeLogger.setGlobalFormatter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLogsBytes_1(t *testing.T) {
	type args struct {
		shardID          int
		cursor           string
		logGroupMaxCount int
	}
	tests := []struct {
		name           string
		s              *LogStore
		args           args
		wantOut        []byte
		wantNextCursor string
		wantErr        bool
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

			gotOut, gotNextCursor, err := tt.s.GetLogsBytes(tt.args.shardID, tt.args.cursor, tt.args.logGroupMaxCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogStore.GetLogsBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("LogStore.GetLogsBytes() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotNextCursor != tt.wantNextCursor {
				t.Errorf("LogStore.GetLogsBytes() gotNextCursor = %v, want %v", gotNextCursor, tt.wantNextCursor)
			}
		})
	}
}

package alils

import (
	"fmt"
	"testing"
)

func TestGetCursor_1(t *testing.T) {
	type args struct {
		shardID int
		from    string
	}
	tests := []struct {
		name       string
		s          *LogStore
		args       args
		wantCursor string
		wantErr    bool
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

			gotCursor, err := tt.s.GetCursor(tt.args.shardID, tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogStore.GetCursor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCursor != tt.wantCursor {
				t.Errorf("LogStore.GetCursor() = %v, want %v", gotCursor, tt.wantCursor)
			}
		})
	}
}

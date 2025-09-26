package session

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestSessionRead_7(t *testing.T) {
	type args struct {
		ctx context.Context
		sid string
	}
	tests := []struct {
		name    string
		pder    *MemProvider
		args    args
		want    Store
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

			got, err := tt.pder.SessionRead(tt.args.ctx, tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemProvider.SessionRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemProvider.SessionRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

package idempotency

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalMsg_2(t *testing.T) {
	type args struct {
		bts []byte
	}
	tests := []struct {
		name    string
		z       *response
		args    args
		wantO   []byte
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

			gotO, err := tt.z.UnmarshalMsg(tt.args.bts)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotO, tt.wantO) {
				t.Errorf("UnmarshalMsg() gotO = %v, want %v", gotO, tt.wantO)
			}
		})
	}
}

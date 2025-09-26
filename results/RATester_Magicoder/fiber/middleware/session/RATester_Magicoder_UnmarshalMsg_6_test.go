package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalMsg_6(t *testing.T) {
	tests := []struct {
		name    string
		z       *data
		bts     []byte
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

			gotO, err := tt.z.UnmarshalMsg(tt.bts)
			if (err != nil) != tt.wantErr {
				t.Errorf("data.UnmarshalMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotO, tt.wantO) {
				t.Errorf("data.UnmarshalMsg() = %v, want %v", gotO, tt.wantO)
			}
		})
	}
}

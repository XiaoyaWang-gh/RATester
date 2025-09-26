package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestdecodeRefArgs_1(t *testing.T) {
	type args struct {
		args map[string]any
	}
	tests := []struct {
		name    string
		p       pageRef
		args    args
		wantRa  refArgs
		wantS   *Site
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

			gotRa, gotS, err := tt.p.decodeRefArgs(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeRefArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRa, tt.wantRa) {
				t.Errorf("decodeRefArgs() gotRa = %v, want %v", gotRa, tt.wantRa)
			}
			if !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("decodeRefArgs() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

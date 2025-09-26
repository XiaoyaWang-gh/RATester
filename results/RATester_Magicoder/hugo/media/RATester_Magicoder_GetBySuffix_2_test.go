package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetBySuffix_2(t *testing.T) {
	type args struct {
		suffix string
	}
	tests := []struct {
		name      string
		t         Types
		args      args
		wantTp    Type
		wantSi    SuffixInfo
		wantFound bool
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

			gotTp, gotSi, gotFound := tt.t.GetBySuffix(tt.args.suffix)
			if !reflect.DeepEqual(gotTp, tt.wantTp) {
				t.Errorf("GetBySuffix() gotTp = %v, want %v", gotTp, tt.wantTp)
			}
			if !reflect.DeepEqual(gotSi, tt.wantSi) {
				t.Errorf("GetBySuffix() gotSi = %v, want %v", gotSi, tt.wantSi)
			}
			if gotFound != tt.wantFound {
				t.Errorf("GetBySuffix() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

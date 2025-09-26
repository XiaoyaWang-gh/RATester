package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetRandomTarget_1(t *testing.T) {
	type fields struct {
		nowIndex   int
		TargetStr  string
		TargetArr  []string
		LocalProxy bool
		RWMutex    sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
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

			s := &Target{
				nowIndex:   tt.fields.nowIndex,
				TargetStr:  tt.fields.TargetStr,
				TargetArr:  tt.fields.TargetArr,
				LocalProxy: tt.fields.LocalProxy,
				RWMutex:    tt.fields.RWMutex,
			}
			got, err := s.GetRandomTarget()
			if (err != nil) != tt.wantErr {
				t.Errorf("Target.GetRandomTarget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Target.GetRandomTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

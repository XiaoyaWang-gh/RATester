package session

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSessionStart_1(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		manager  *Manager
		args     args
		wantErr  bool
		wantSess bool
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

			gotSession, err := tt.manager.SessionStart(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.SessionStart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (gotSession != nil) != tt.wantSess {
				t.Errorf("Manager.SessionStart() = %v, want %v", gotSession, tt.wantSess)
			}
		})
	}
}

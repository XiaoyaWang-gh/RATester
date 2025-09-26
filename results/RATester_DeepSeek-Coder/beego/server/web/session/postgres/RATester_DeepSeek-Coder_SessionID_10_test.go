package postgres

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_10(t *testing.T) {
	type fields struct {
		sid string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
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

			st := &SessionStore{
				sid: tt.fields.sid,
			}
			if got := st.SessionID(tt.args.ctx); got != tt.want {
				t.Errorf("SessionID() = %v, want %v", got, tt.want)
			}
		})
	}
}

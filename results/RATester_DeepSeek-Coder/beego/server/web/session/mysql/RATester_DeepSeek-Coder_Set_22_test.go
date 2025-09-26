package mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestSet_22(t *testing.T) {
	type args struct {
		ctx   context.Context
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name    string
		st      *SessionStore
		args    args
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

			st := &SessionStore{
				c:      tt.st.c,
				sid:    tt.st.sid,
				values: tt.st.values,
			}
			if err := st.Set(tt.args.ctx, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("SessionStore.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

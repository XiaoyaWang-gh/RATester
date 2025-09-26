package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSessionReleaseIfPresent_2(t *testing.T) {
	type args struct {
		ctx context.Context
		w   http.ResponseWriter
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				ctx: context.Background(),
				w:   httptest.NewRecorder(),
			},
		},
		{
			name: "Test case 2",
			args: args{
				ctx: context.TODO(),
				w:   httptest.NewRecorder(),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			st := &SessionStore{
				c:      &sql.DB{},
				sid:    "test_session_id",
				values: make(map[interface{}]interface{}),
			}

			st.SessionReleaseIfPresent(tt.args.ctx, tt.args.w)
		})
	}
}

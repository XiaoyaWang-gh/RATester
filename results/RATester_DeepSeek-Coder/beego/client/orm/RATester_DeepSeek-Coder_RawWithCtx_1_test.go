package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestRawWithCtx_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		query   string
		args    []interface{}
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "valid query",
			query:   "SELECT * FROM users WHERE id = ?",
			args:    []interface{}{1},
			wantErr: false,
		},
		{
			name:    "invalid query",
			query:   "SELECT * FROM users WHERE id = ?",
			args:    []interface{}{"1"},
			wantErr: true,
		},
		{
			name:    "empty query",
			query:   "",
			args:    []interface{}{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := context.Background()
			o := &ormBase{}
			_, err := o.RawWithCtx(ctx, tc.query, tc.args...).Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("RawWithCtx() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

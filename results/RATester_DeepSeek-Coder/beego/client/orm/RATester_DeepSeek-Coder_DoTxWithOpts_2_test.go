package orm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
)

func TestDoTxWithOpts_2(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		opts    *sql.TxOptions
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "success",
			opts:    &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false},
			wantErr: false,
		},
		{
			name:    "failure",
			opts:    &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: true},
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

			d := &DoNothingOrm{}
			err := d.DoTxWithOpts(tc.opts, func(ctx context.Context, txOrm TxOrmer) error {
				if tc.wantErr {
					return errors.New("mock error")
				}
				return nil
			})

			if (err != nil) != tc.wantErr {
				t.Errorf("DoTxWithOpts() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

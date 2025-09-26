package orm

import (
	"fmt"
	"testing"
)

func TestRaw_4(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		query   string
		args    []interface{}
		wantErr bool
	}{
		{
			name:    "success",
			query:   "SELECT * FROM table",
			args:    []interface{}{},
			wantErr: false,
		},
		{
			name:    "failure",
			query:   "SELECT * FROM table WHERE id = ?",
			args:    []interface{}{1},
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

			ormer := &filterOrmDecorator{}
			rawSeter := ormer.Raw(tc.query, tc.args...)
			_, err := rawSeter.Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("Raw() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

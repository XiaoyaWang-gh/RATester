package orm

import (
	"fmt"
	"testing"
)

func TestRowsToStruct_3(t *testing.T) {
	type TestStruct struct {
		Key   string
		Value string
	}

	testCases := []struct {
		name     string
		query    string
		args     []interface{}
		keyCol   string
		valueCol string
		want     int64
		wantErr  bool
	}{
		{
			name:     "success",
			query:    "SELECT key, value FROM test_table",
			args:     []interface{}{},
			keyCol:   "key",
			valueCol: "value",
			want:     10,
			wantErr:  false,
		},
		{
			name:     "failure",
			query:    "SELECT key, value FROM test_table",
			args:     []interface{}{},
			keyCol:   "key",
			valueCol: "value",
			want:     0,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			o := &rawSet{
				query: tc.query,
				args:  tc.args,
			}

			got, err := o.RowsToStruct(TestStruct{}, tc.keyCol, tc.valueCol)
			if (err != nil) != tc.wantErr {
				t.Errorf("RowsToStruct() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("RowsToStruct() = %v, want %v", got, tc.want)
			}
		})
	}
}

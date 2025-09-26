package orm

import (
	"fmt"
	"testing"
)

func TestMaxStmtCacheSize_1(t *testing.T) {
	testCases := []struct {
		name string
		arg  int
		want int
	}{
		{
			name: "Test case 1",
			arg:  10,
			want: 10,
		},
		{
			name: "Test case 2",
			arg:  20,
			want: 20,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			option := MaxStmtCacheSize(tc.arg)
			alias := &alias{}
			option(alias)
			if alias.StmtCacheSize != tc.want {
				t.Errorf("MaxStmtCacheSize() = %v, want %v", alias.StmtCacheSize, tc.want)
			}
		})
	}
}

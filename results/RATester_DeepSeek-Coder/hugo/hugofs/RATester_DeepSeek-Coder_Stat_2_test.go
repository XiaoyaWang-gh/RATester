package hugofs

import (
	"fmt"
	"testing"
)

func TestStat_2(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		setup   func(t *testing.T) *rootMappingDir
		wantErr error
	}{
		{
			name: "should return errIsDir when Stat is called",
			setup: func(t *testing.T) *rootMappingDir {
				return &rootMappingDir{
					name: "test",
					meta: &FileMeta{},
				}
			},
			wantErr: errIsDir,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			dir := tc.setup(t)
			_, err := dir.Stat()
			if err != tc.wantErr {
				t.Errorf("expected error %v, got %v", tc.wantErr, err)
			}
		})
	}
}

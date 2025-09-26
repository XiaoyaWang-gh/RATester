package web

import (
	"fmt"
	"testing"
)

func TestAddViewPath_1(t *testing.T) {
	testCases := []struct {
		name     string
		viewPath string
		wantErr  bool
	}{
		{
			name:     "Test case 1",
			viewPath: "path/to/views",
			wantErr:  false,
		},
		{
			name:     "Test case 2",
			viewPath: "path/to/other/views",
			wantErr:  false,
		},
		{
			name:     "Test case 3",
			viewPath: "path/to/another/views",
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

			err := AddViewPath(tc.viewPath)
			if (err != nil) != tc.wantErr {
				t.Errorf("AddViewPath() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

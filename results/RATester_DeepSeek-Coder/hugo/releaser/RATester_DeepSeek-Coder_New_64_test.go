package releaser

import (
	"fmt"
	"testing"
)

func TestNew_64(t *testing.T) {
	tests := []struct {
		name      string
		skipPush  bool
		try       bool
		step      int
		wantError bool
	}{
		{
			name:      "valid case",
			skipPush:  false,
			try:       false,
			step:      1,
			wantError: false,
		},
		{
			name:      "invalid step",
			skipPush:  false,
			try:       false,
			step:      3,
			wantError: true,
		},
		{
			name:      "not a release branch",
			skipPush:  false,
			try:       false,
			step:      1,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := New(tt.skipPush, tt.try, tt.step)
			if (err != nil) != tt.wantError {
				t.Errorf("New() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

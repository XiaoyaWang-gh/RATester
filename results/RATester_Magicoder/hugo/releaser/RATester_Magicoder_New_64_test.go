package releaser

import (
	"fmt"
	"testing"
)

func TestNew_64(t *testing.T) {
	tests := []struct {
		name     string
		skipPush bool
		try      bool
		step     int
		wantErr  bool
	}{
		{
			name:     "Test case 1",
			skipPush: false,
			try:      false,
			step:     1,
			wantErr:  false,
		},
		{
			name:     "Test case 2",
			skipPush: true,
			try:      true,
			step:     2,
			wantErr:  false,
		},
		{
			name:     "Test case 3",
			skipPush: false,
			try:      false,
			step:     3,
			wantErr:  true,
		},
		{
			name:     "Test case 4",
			skipPush: true,
			try:      true,
			step:     0,
			wantErr:  true,
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
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

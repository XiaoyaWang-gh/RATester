package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat_6(t *testing.T) {
	testCases := []struct {
		name string
		lm   *LogMsg
		want string
	}{
		{
			name: "TestFormat_ErrorLevel",
			lm: &LogMsg{
				Level:      LevelError,
				Msg:        "This is an error message",
				When:       time.Now(),
				FilePath:   "test.go",
				LineNumber: 10,
			},
			want: "test.go:10 [ERRO] This is an error message",
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

			m := &multiFileLogWriter{}
			got := m.Format(tc.lm)
			if got != tc.want {
				t.Errorf("Format() = %v, want %v", got, tc.want)
			}
		})
	}
}

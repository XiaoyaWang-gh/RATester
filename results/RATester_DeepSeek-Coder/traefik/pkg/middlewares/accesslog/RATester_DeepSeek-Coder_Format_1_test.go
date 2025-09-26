package accesslog

import (
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestFormat_1(t *testing.T) {
	testCases := []struct {
		name     string
		entry    *logrus.Entry
		expected string
	}{
		{
			name: "test case 1",
			entry: &logrus.Entry{
				Data: logrus.Fields{
					ClientHost:            "192.0.2.0",
					ClientUsername:        "jdoe",
					StartUTC:              time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					RequestMethod:         "GET",
					RequestPath:           "/path",
					RequestProtocol:       "HTTP/1.1",
					DownstreamStatus:      200,
					DownstreamContentSize: 1024,
					Duration:              time.Duration(100) * time.Millisecond,
				},
			},
			expected: "192.0.2.0 - jdoe [01/Jan/2022:00:00:00 +0000] \"GET /path HTTP/1.1\" 200 1024 100ms\n",
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

			formatter := &CommonLogFormatter{}
			result, err := formatter.Format(tc.entry)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if string(result) != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, string(result))
			}
		})
	}
}

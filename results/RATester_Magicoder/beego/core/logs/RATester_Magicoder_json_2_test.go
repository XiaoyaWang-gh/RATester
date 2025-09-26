package logs

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Testjson_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	record := &AccessLogRecord{
		RemoteAddr:     "127.0.0.1",
		RequestTime:    time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		RequestMethod:  "GET",
		Request:        "/",
		ServerProtocol: "HTTP/1.1",
		Host:           "localhost",
		Status:         200,
		BodyBytesSent:  1024,
		ElapsedTime:    time.Second,
		HTTPReferrer:   "http://example.com",
		HTTPUserAgent:  "Mozilla/5.0",
		RemoteUser:     "user",
	}

	jsonBytes, err := record.json()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	var decodedRecord AccessLogRecord
	err = json.Unmarshal(jsonBytes, &decodedRecord)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(record, &decodedRecord) {
		t.Fatalf("Expected: %v, got: %v", record, &decodedRecord)
	}
}

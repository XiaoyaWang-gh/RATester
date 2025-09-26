package accesslog

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetLogData_1(t *testing.T) {
	t.Run("TestGetLogData", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		req, err := http.NewRequest("GET", "/test", nil)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.WithValue(req.Context(), DataTableKey, &LogData{})
		req = req.WithContext(ctx)

		ld := GetLogData(req)
		if ld == nil {
			t.Fatal("expected LogData, got nil")
		}
	})
}

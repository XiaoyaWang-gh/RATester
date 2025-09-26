package middleware

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestLimitedReaderPool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := BodyLimitConfig{
		Skipper: func(c echo.Context) bool {
			return false
		},
		Limit: "100B",
		limit: 100,
	}

	pool := limitedReaderPool(config)

	obj := pool.Get().(*limitedReader)
	if obj.Limit != config.Limit {
		t.Errorf("Expected Limit to be %s, but got %s", config.Limit, obj.Limit)
	}

	pool.Put(obj)

	obj2 := pool.Get().(*limitedReader)
	if obj2.Limit != config.Limit {
		t.Errorf("Expected Limit to be %s, but got %s", config.Limit, obj2.Limit)
	}
}

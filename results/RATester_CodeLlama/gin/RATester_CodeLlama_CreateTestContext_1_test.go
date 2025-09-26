package gin

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestCreateTestContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := httptest.NewRecorder()
	c, r := CreateTestContext(w)
	if c == nil || r == nil {
		t.Errorf("CreateTestContext() failed")
	}
}

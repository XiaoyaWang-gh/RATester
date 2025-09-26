package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetSameSite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Params: Params{},
	}
	c.SetSameSite(http.SameSiteStrictMode)
	if c.sameSite != http.SameSiteStrictMode {
		t.Errorf("SetSameSite() = %v, want %v", c.sameSite, http.SameSiteStrictMode)
	}
}

package apiauth

import (
	"fmt"
	"testing"
)

func TestAPIBasicAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	appid := "appid"
	appkey := "appkey"
	f := APIBasicAuth(appid, appkey)
	if f == nil {
		t.Errorf("APIBasicAuth() = %v, want %v", f, nil)
	}
}

package fiber

import (
	"fmt"
	"testing"
)

func TestParamsMatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	specParamStr := headerParams{"key1": []byte("value1"), "key2": []byte("value2")}
	offerParams := "key1=value1,key2=value2"
	if !paramsMatch(specParamStr, offerParams) {
		t.Errorf("paramsMatch failed")
	}
}

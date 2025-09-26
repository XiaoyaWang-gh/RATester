package proxy

import (
	"fmt"
	"testing"
)

func TestClose_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	httpsListener := &HttpsListener{}
	err := httpsListener.Close()
	if err != nil {
		t.Errorf("Close() error = %v, want nil", err)
		return
	}
}

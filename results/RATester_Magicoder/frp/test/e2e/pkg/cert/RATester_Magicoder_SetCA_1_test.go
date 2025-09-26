package cert

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetCA_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cp := &SelfSignedCertGenerator{}
	caKey := []byte("testKey")
	caCert := []byte("testCert")
	cp.SetCA(caKey, caCert)

	if !reflect.DeepEqual(cp.caKey, caKey) || !reflect.DeepEqual(cp.caCert, caCert) {
		t.Errorf("Expected caKey and caCert to be set correctly, but got %v and %v", cp.caKey, cp.caCert)
	}
}

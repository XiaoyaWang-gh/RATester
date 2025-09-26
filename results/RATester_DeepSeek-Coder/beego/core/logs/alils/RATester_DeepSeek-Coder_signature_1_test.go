package alils

import (
	"fmt"
	"testing"
)

func TestSignature_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	project := &LogProject{
		Name:            "test",
		Endpoint:        "http://localhost:8080",
		AccessKeyID:     "test",
		AccessKeySecret: "test",
	}

	headers := map[string]string{
		"Content-MD5":   "test",
		"Content-Type":  "test",
		"Date":          "test",
		"X-Sls-Header1": "test",
		"X-Sls-Header2": "test",
	}

	digest, err := signature(project, "GET", "http://localhost:8080/test", headers)
	if err != nil {
		t.Errorf("signature() error = %v", err)
		return
	}

	if digest == "" {
		t.Errorf("signature() = %v, want non-empty string", digest)
	}
}

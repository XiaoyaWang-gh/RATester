package passtlsclientcert

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	passTLSClientCert := &passTLSClientCert{
		name: "testName",
	}
	name, typeName, spanKind := passTLSClientCert.GetTracingInformation()

	if name != "testName" {
		t.Errorf("Expected name to be 'testName', but got %s", name)
	}

	if typeName != "passTLSClientCert" {
		t.Errorf("Expected typeName to be 'passTLSClientCert', but got %s", typeName)
	}

	if spanKind != trace.SpanKindInternal {
		t.Errorf("Expected spanKind to be trace.SpanKindInternal, but got %v", spanKind)
	}
}

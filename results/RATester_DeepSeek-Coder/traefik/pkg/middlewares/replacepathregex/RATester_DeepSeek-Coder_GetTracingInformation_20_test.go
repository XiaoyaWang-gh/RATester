package replacepathregex

import (
	"fmt"
	"regexp"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rp := &replacePathRegex{
		next:        nil,
		regexp:      regexp.MustCompile("^/api/v1/"),
		replacement: "/v1/",
		name:        "replacePathRegex",
	}

	name, typeName, spanKind := rp.GetTracingInformation()

	if name != "replacePathRegex" {
		t.Errorf("Expected name to be 'replacePathRegex', got %s", name)
	}

	if typeName != "replacePathRegex" {
		t.Errorf("Expected typeName to be 'replacePathRegex', got %s", typeName)
	}

	if spanKind != trace.SpanKindInternal {
		t.Errorf("Expected spanKind to be trace.SpanKindInternal, got %v", spanKind)
	}
}

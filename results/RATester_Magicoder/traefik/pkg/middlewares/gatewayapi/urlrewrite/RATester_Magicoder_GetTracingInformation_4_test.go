package urlrewrite

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	u := urlRewrite{
		name: "test",
	}
	name, typeName, kind := u.GetTracingInformation()
	if name != "test" {
		t.Errorf("Expected name to be 'test', but got %s", name)
	}
	if typeName != "urlRewrite" {
		t.Errorf("Expected typeName to be 'urlRewrite', but got %s", typeName)
	}
	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %v", kind)
	}
}

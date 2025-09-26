package postgres

import (
	"context"
	"fmt"
	"testing"
)

func TestGet_30(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	st := &SessionStore{}
	st.values = make(map[interface{}]interface{})
	st.values["key"] = "value"
	ctx := context.Background()
	if st.Get(ctx, "key") != "value" {
		t.Errorf("st.Get(ctx, \"key\") = %v, want \"value\"", st.Get(ctx, "key"))
	}
}

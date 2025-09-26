package binding

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBind_9(t *testing.T) {
	hb := headerBinding{}

	type testStruct struct {
		Header http.Header
		Obj    any
	}

	tests := []struct {
		name    string
		args    testStruct
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: testStruct{
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Obj: &struct{}{},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: testStruct{
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Obj: "invalid",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := hb.Bind(&http.Request{Header: tt.args.Header}, tt.args.Obj); (err != nil) != tt.wantErr {
				t.Errorf("headerBinding.Bind() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

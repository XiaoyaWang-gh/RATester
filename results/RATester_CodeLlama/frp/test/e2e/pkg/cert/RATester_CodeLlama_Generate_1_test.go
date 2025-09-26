package cert

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		cp   *SelfSignedCertGenerator
		want *Artifacts
	}{
		{
			name: "valid",
			cp: &SelfSignedCertGenerator{
				caKey:  []byte("ca-key"),
				caCert: []byte("ca-cert"),
			},
			want: &Artifacts{
				Key:    []byte("key"),
				Cert:   []byte("cert"),
				CAKey:  []byte("ca-key"),
				CACert: []byte("ca-cert"),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			got, err := tt.cp.Generate("localhost")
			if err != nil {
				t.Fatalf("failed to generate certs: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

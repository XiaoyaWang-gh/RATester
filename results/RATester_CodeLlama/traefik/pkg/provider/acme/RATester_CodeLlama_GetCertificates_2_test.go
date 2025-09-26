package acme

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestGetCertificates_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	storedData := &StoredData{
		Account: &Account{},
		Certificates: []*CertAndStore{
			{
				Certificate: Certificate{
					Domain: types.Domain{
						Main: "test.com",
					},
				},
			},
		},
	}

	s := &LocalStore{
		storedData: map[string]*StoredData{
			"test": storedData,
		},
	}

	// when
	certificates, err := s.GetCertificates("test")

	// then
	require.NoError(t, err)
	require.Len(t, certificates, 1)
	require.Equal(t, storedData.Certificates[0], certificates[0])
}

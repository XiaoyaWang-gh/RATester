package tls

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestGet_4(t *testing.T) {
	manager := &Manager{
		lock:         sync.RWMutex{},
		storesConfig: map[string]Store{},
		stores:       map[string]*CertificateStore{},
		configs:      map[string]Options{},
		certs:        []*CertAndStores{},
	}

	tests := []struct {
		name       string
		storeName  string
		configName string
		want       *tls.Config
		wantErr    bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := manager.Get(tt.storeName, tt.configName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

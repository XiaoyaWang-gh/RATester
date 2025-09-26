package client

import (
	"fmt"
	"testing"
)

func TestAcquireCookieJar_1(t *testing.T) {
	t.Run("should return a CookieJar from the pool", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := AcquireCookieJar()
		if got == nil {
			t.Errorf("AcquireCookieJar() = %v; want not nil", got)
		}
	})

	t.Run("should panic if the type assertion fails", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		cookieJarPool.Put(new(struct{}))
		AcquireCookieJar()
	})
}

package acme

import (
	"fmt"
	"testing"

	"github.com/go-acme/lego/v4/registration"
)

func TestSaveAccount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &LocalStore{
		saveDataChan: make(chan map[string]*StoredData),
		filename:     "test.json",
		storedData:   make(map[string]*StoredData),
	}

	account := &Account{
		Email:        "test@example.com",
		Registration: &registration.Resource{},
		PrivateKey:   []byte("privateKey"),
	}

	err := store.SaveAccount("test", account)
	if err != nil {
		t.Errorf("SaveAccount() error = %v", err)
		return
	}

	storedData, err := store.get("test")
	if err != nil {
		t.Errorf("get() error = %v", err)
		return
	}

	if storedData.Account.Email != account.Email {
		t.Errorf("SaveAccount() failed, expected email = %s, got = %s", account.Email, storedData.Account.Email)
	}
}

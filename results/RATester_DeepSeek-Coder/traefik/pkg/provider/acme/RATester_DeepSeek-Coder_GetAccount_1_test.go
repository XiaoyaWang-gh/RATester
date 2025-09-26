package acme

import (
	"fmt"
	"testing"
)

func TestGetAccount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &LocalStore{
		storedData: make(map[string]*StoredData),
	}

	testAccount := &Account{
		Email: "test@example.com",
	}

	testResolverName := "testResolver"

	store.storedData[testResolverName] = &StoredData{
		Account: testAccount,
	}

	account, err := store.GetAccount(testResolverName)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if account.Email != testAccount.Email {
		t.Errorf("Expected account email to be %s, got %s", testAccount.Email, account.Email)
	}
}

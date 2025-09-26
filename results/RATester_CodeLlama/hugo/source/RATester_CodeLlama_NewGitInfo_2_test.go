package source

import (
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/bep/gitmap"
)

func TestNewGitInfo_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	info := gitmap.GitInfo{
		Hash:            "1234567890",
		AbbreviatedHash: "1234567",
		Subject:         "This is a commit message",
		AuthorName:      "John Doe",
		AuthorEmail:     "john.doe@example.com",
		AuthorDate:      time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		CommitDate:      time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		Body:            "This is the commit message body",
	}

	// Act
	result := NewGitInfo(info)

	// Assert
	assert.Equal(t, "1234567890", result.Hash)
	assert.Equal(t, "1234567", result.AbbreviatedHash)
	assert.Equal(t, "This is a commit message", result.Subject)
	assert.Equal(t, "John Doe", result.AuthorName)
	assert.Equal(t, "john.doe@example.com", result.AuthorEmail)
	assert.Equal(t, time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC), result.AuthorDate)
	assert.Equal(t, time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC), result.CommitDate)
	assert.Equal(t, "This is the commit message body", result.Body)
}

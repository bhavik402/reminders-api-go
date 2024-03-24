package reminders_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/bhavik402/reminders-api-go/api-rest/internal/data/reminders"
	"github.com/stretchr/testify/assert"
)

var (
	testDBFile = "testDB.json"
)

func NewMockJsonStorage() (*reminders.JsonStorage, error) {
	return &reminders.JsonStorage{
		FileName: testDBFile,
	}, nil
}

func setup() *reminders.JsonStorage {
	s, err := NewMockJsonStorage()
	if err != nil {
		fmt.Printf("failed to create NewMockJsonStorage: %s", err)
	}
	return s
}

func cleanup() {
	err := os.Remove(testDBFile)
	if err != nil {
		fmt.Printf("failed to delete test json file %s", err)
	}
}

func TestSimpleWritet(t *testing.T) {
	s := setup()
	r := reminders.Reminder{
		Title: "test title",
	}

	err := s.Save(&r)
	if err != nil {
		t.Error("failed to save reminder: %w", err)
	}

	assert.Equal(t, nil, err, "err is nil")
}

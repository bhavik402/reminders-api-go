package reminders_test

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bhavik402/reminders-api-go/api-rest/internal/data/reminders"
)

var (
	testDBFile = "testDB.json"
	mockDB     = `[{"title":"title 1","status":"PENDING","notes":"notes for this reminder 1","category":"chores","priority":"LOW","flag":false,"tags":["tag1","tag2","tag3"]},{"title":"title 2","status":"COMPLETED","notes":"notes for this reminder 1","category":"chores","priority":"HIGH","flag":false,"tags":["tag1","tag2","tag3"],"dueOn":"2024-03-29T13:11:35.313Z"},{"title":"title 3","status":"COMPLETED","notes":"notes for this reminder 1","category":"chores","priority":"HIGH","flag":true,"tags":["tag1","tag2","tag3"]},{"title":"title 4","status":"PENDING","notes":"notes for this reminder 1","category":"chores","priority":"LOW","flag":false,"tags":["tag1","tag2","tag3"],"dueOn":"2024-04-29T13:11:35.313Z"}]`
)

func NewMockJsonStorage() (*reminders.JsonStorage, error) {
	return &reminders.JsonStorage{
		FileName: testDBFile,
	}, nil
}

func setupNewMockStorage() (*reminders.JsonStorage, error) {
	s, err := NewMockJsonStorage()
	if err != nil {
		return nil, fmt.Errorf("failed to create NewMockJsonStorage: %w", err)
	}
	return s, nil
}

// Load data from mock string and save it as JSON file for "db" usage
func setupJsonDBFile() (*reminders.JsonStorage, error) {
	s, err := setupNewMockStorage()
	if err != nil {
		return nil, err
	}

	var reminders []reminders.Reminder
	err = json.Unmarshal([]byte(mockDB), &reminders)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal mock data: %w", err)
	}

	for _, v := range reminders {
		err = s.Save(&v)
		if err != nil {
			return nil, fmt.Errorf("failed to save reminder: %w", err)
		}
	}

	return s, nil
}

// Delete the JSON db file once tests are complete
func cleanup() {
	err := os.Remove(testDBFile)
	if err != nil {
		fmt.Printf("failed to delete test json file %s", err)
	}
}

// Tests saving the data to Json file
func TestCreatingJsonDBFile(t *testing.T) {
	_, err := setupJsonDBFile()
	if err != nil {
		t.Errorf("failed to setup db and write to file")
	}

	assert.Equal(t, nil, err, "recieved an error while saving a file")
	cleanup()
}

// Tests saving individual record but also not loosing the existing
func TestRecordInsertion(t *testing.T) {
	s, err := setupJsonDBFile()
	if err != nil {
		t.Errorf("failed to setup db and write to file")
	}

	r := reminders.Reminder{
		Title:    "dummy Title",
		Status:   reminders.PENDING,
		Notes:    "dummy notes",
		Category: "category",
		Flag:     true,
	}

	remsBefore, err := s.ReadAll()
	if err != nil {
		t.Errorf("failed to read db")
	}

	err = s.Save(&r)

	allReminders, err := s.ReadAll()
	if err != nil {
		t.Errorf("failed to read db")
	}

	gotIndex := slices.IndexFunc(allReminders, func(rem reminders.Reminder) bool { return rem.Title == r.Title })

	assert.Equal(t, nil, err, "recieved an error while saving a file")
	cleanup()
	assert.Equal(t, r.Title, allReminders[gotIndex].Title)
	assert.Equal(t, r.Status, allReminders[gotIndex].Status)
	assert.Equal(t, r.Notes, allReminders[gotIndex].Notes)
	assert.Equal(t, r.Category, allReminders[gotIndex].Category)
	assert.Equal(t, r.Flag, allReminders[gotIndex].Flag)
	assert.Equal(t, len(remsBefore)+1, len(allReminders))
}

// Read all data out of json file
// Read individual data out of json file
// Change the status of the reminder
// Change the flag of the reminder
// Remove a specific reminder
// Test all the different types of errors that could occur.
// Errors tests for ErrFailedToOpenFile, ErrFailedToReadFile

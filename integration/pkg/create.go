package pkg

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
)

//go:embed reminders_records.json
var remindersRecords []byte

type Reminder struct {
	Title    string   `json:"title,omitempty"`
	Status   string   `json:"status,omitempty"`
	Notes    string   `json:"notes,omitempty"`
	Category string   `json:"category,omitempty"`
	Priority string   `json:"priority,omitempty"`
	Flag     bool     `json:"flag,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	DueOn    string   `json:"dueOn,omitempty"`
}

func ReadRecords() ([]Reminder, error) {
	var rr []Reminder
	err := json.Unmarshal(remindersRecords, &rr)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall reminder records: %w", err)
	}

	return rr, nil
}

func PostAllRecords() error {
	var url string = "http://localhost:8452/v1/reminders"

	records, err := ReadRecords()
	if err != nil {
		return err
	}

	for _, v := range records {
		v, _ := json.Marshal(v)
		rb := bytes.NewReader(v)

		res, err := http.Post(url, "application/json", rb)
		if err != nil {
			return err
		}

		if res.StatusCode != http.StatusOK {
			fmt.Println("❌ Failed to Post")
		}
	}

	fmt.Printf("✅ Test for %s Successfully completed \n", url)

	return nil
}

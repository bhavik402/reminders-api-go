package pkg

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pterm/pterm"
)

//go:embed reminders_records.json
var remindersRecords []byte

func ReadRecords() ([]Reminder, error) {
	var rr []Reminder
	err := json.Unmarshal(remindersRecords, &rr)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall reminder records: %w", err)
	}

	return rr, nil
}

func RunAllPostTests(logger *pterm.Logger) error {
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
			return fmt.Errorf("❌ Failed to POST record: with StatusCode == %s: ", strconv.Itoa(res.StatusCode))
		}
	}

	logOutput := fmt.Sprintf("✅ Test for POST: %s Successfully Completed \n", url)
	logger.Info(logOutput)

	return nil
}

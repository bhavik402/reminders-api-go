package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/pterm/pterm"
)

func RunAllGetTests(logger *pterm.Logger) error {
	var url string = "http://localhost:8452/v1/reminders"

	records, err := requestAndGetAllRecords(url)
	if err != nil {
		return err
	}

	if len(records) <= 0 {
		return fmt.Errorf("❌ Failed to GET any records")
	}

	logOutput := fmt.Sprintf("✅ Test for GET: %s Successfully Completed with %s many records \n", url, strconv.Itoa(len(records)))
	logger.Info(logOutput)

	return nil
}

func requestAndGetAllRecords(url string) ([]Reminder, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("❌ Failed to GET records: with StatusCode == %s: ", strconv.Itoa(res.StatusCode))
	}

	records := make([]Reminder, 0)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

package pkg

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
)

func RunAllPutTests() (*string, error) {
	var url string = "http://localhost:8452/v1/reminders"
	records, err := requestAndGetAllRecords(url)
	if err != nil {
		return nil, err
	}

	i := slices.IndexFunc(records, func(r Reminder) bool { return r.Flag })
	if i == -1 {
		return nil, fmt.Errorf("failed to find a record with flag: true")
	}
	id := records[i].Id

	//make a PUT Request
	putFlagUrl := fmt.Sprintf("%s/flag/%s", url, id)
	req, err := http.NewRequest(http.MethodPut, putFlagUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update flag: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request to update flag: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update flag, got Response Status Code: %s", strconv.Itoa(resp.StatusCode))
	}

	//make a GET Request

	result := fmt.Sprintf("✅ Test for PUT Flag: %s Successfully Completed\n", putFlagUrl)

	return &result, nil
}

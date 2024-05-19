package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Currency struct {
	Rate float64 `json:"rate"`
}

func GetRate() (float64, error) {
	client := &http.Client{}

	now := time.Now()
	today := now.Format("20060102")

	url := fmt.Sprintf("https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?valcode=USD&date=%s&json", today)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %v", err)
	}

	var currencies []Currency
	if err := json.Unmarshal(body, &currencies); err != nil {
		return 0, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	if len(currencies) == 0 {
		return 0, fmt.Errorf("no currency data available")
	}

	return currencies[0].Rate, nil
}

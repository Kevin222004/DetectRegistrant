package main

import (
	"encoding/csv"
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"
)

type Data struct {
	Registrant       string
	ModelDisplayName string
	Model            string
	SourceURL        string
}

type Sheet struct {
	Data []*Data
}

func parseCSV(filename string) (*Sheet, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}

	sheet := &Sheet{
		Data: make([]*Data, 0, len(records)-1),
	}

	// Skip the header row
	for i, record := range records[1:] {
		if len(record) < 4 {
			return nil, fmt.Errorf("row %d has insufficient columns", i+2)
		}

		data := &Data{
			Registrant:       strings.TrimSpace(record[0]),
			ModelDisplayName: strings.TrimSpace(record[1]),
			Model:            strings.TrimSpace(record[2]),
			SourceURL:        strings.TrimSpace(record[3]),
		}
		sheet.Data = append(sheet.Data, data)
	}

	return sheet, nil
}

func TestFindRegistrant(t *testing.T) {
	sheet, err := parseCSV("test.csv")
	if err != nil {
		t.Fatalf("Failed to parse CSV: %v", err)
	}

	totalModel := 0
	for _, data := range sheet.Data {

		parsedURL, err := url.Parse(data.SourceURL)
		if err != nil {
			t.Errorf("Failed to parse URL %s: %v", data.SourceURL, err)
			continue
		}

		result := FindRegistrant(parsedURL)

		if strings.ToLower(result) != strings.ToLower(data.Registrant) {
			t.Errorf("FindRegistrant(%s) = %s; expected is  %s ModelDisplayName: %s",
				data.SourceURL, result, data.Registrant, data.ModelDisplayName)
		}

		totalModel++
	}

	fmt.Println("checked ", totalModel)
}

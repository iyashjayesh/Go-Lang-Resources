package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type LogLine struct {
	Timestamp int64  `json:"timestamp"`
	Line      string `json:"line"`
	File      string `json:"file"`
}

type LogRequest struct {
	Lines []LogLine `json:"lines"`
}

func sendLogToLogDNA(apiKey, hostname, mac, ip, logLine, logFile string) error {
	logData := LogRequest{
		Lines: []LogLine{
			{
				Timestamp: time.Now().UnixMilli(), // Current timestamp in milliseconds
				Line:      logLine,
				File:      logFile,
			},
		},
	}
	jsonData, err := json.Marshal(logData)
	if err != nil {
		return fmt.Errorf("failed to marshal log data: %v", err)
	}

	url := fmt.Sprintf("https://logs.logdna.com/logs/ingest?hostname=%s&mac=%s&ip=%s&now=%d", hostname, mac, ip, time.Now().UnixMilli())
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.SetBasicAuth(apiKey, "") // API key with no password

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send log to LogDNA: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("received non-200 response: %s", resp.Status)
	}

	log.Println("Log successfully sent to LogDNA")
	return nil
}

func main() {
	apiKey := "API_KEY"
	hostname := "EXAMPLE_HOST"
	mac := "C0:FF:EE:C0:FF:EE"
	ip := "10.0.1.101"
	logLine := "This is an awesome log statement"
	logFile := "example.log"

	err := sendLogToLogDNA(apiKey, hostname, mac, ip, logLine, logFile)
	if err != nil {
		fmt.Printf("Error sending log: %v\n", err)
	}
}

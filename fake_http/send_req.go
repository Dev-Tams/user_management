package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Dev-Tams/user_management/user"
)

func GetReq(url string) (string, error) {

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return "", err

	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(" bad status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body: %w", err)
	}

	return string(body), nil

}

func PostReq(url string, p user.Post) (*user.CreatePost, error) {


	payload, err := json.Marshal(p)
	if err != nil{
		return nil, fmt.Errorf(" Error marshalling payload %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://jsonplaceholder.typicode.com/posts", bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf(" bad status %v", req.Response)
	}

	// set header
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(" Timeout error %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf(" bad status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	var createdPost user.CreatePost
	if err := json.Unmarshal(body, &createdPost); err != nil {
		return nil, fmt.Errorf("unmarshalling response body: %w", err)
	}

	return &createdPost, nil
}

package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Dev-Tams/user_management/user"
)

func GetReq(ctx context.Context, url string) (string, error) {

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := client.Do(req)
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

func PostReq(ctx context.Context, url string, p user.Post) (*user.CreatePost, error) {


	payload, err := json.Marshal(p)
	if err != nil{
		return nil, fmt.Errorf(" Error marshalling payload %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// set header
	req.Header.Set("Content-Type", "application/json")

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

package request

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func SendReq(url string) (string, error) {

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

	body, err:= io.ReadAll(resp.Body)
	if err != nil{
		return "", fmt.Errorf("reading response body: %w", err)
	}
	

	return string(body), nil

}

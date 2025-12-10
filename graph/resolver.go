package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	FollowServiceURL string
	HTTPClient       *http.Client
}

func (r *Resolver) callFollowService(method, endpoint string, payload interface{}, result interface{}) error {
	var body io.Reader
	if payload != nil {
		b, _ := json.Marshal(payload)
		body = bytes.NewBuffer(b)
	}

	url := fmt.Sprintf("%s%s", r.FollowServiceURL, endpoint)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := r.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errResp map[string]string
		json.NewDecoder(resp.Body).Decode(&errResp)
		return fmt.Errorf(errResp["error"])
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

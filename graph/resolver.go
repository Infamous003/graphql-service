package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Infamous003/graphql-service/graph/model"
)

type Resolver struct {
	FollowServiceURL string
	HTTPClient       *http.Client
}

func (r *Resolver) callFollowService(method, endpoint string, payload any, result any) error {
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
		return fmt.Errorf("%v", errResp["error"])
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func convertUsers(users []*model.User) []*model.User {
	result := make([]*model.User, len(users))
	for i, u := range users {
		result[i] = &model.User{
			ID:        u.ID,
			Username:  u.Username,
			CreatedAt: u.CreatedAt,
		}
	}
	return result
}

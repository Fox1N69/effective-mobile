package infra

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test-task/internal/models"
	"time"

	"github.com/pkg/errors"
)

type UserAPIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewUserAPIClient(baseURL string) *UserAPIClient {
	return &UserAPIClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *UserAPIClient) FetchUserInfo(passportSerie, passportNumber string) (*models.UserInfo, error) {
	reqURL := fmt.Sprintf("%s/info?passportSerie=%s&passportNumber=%s", c.BaseURL, passportSerie, passportNumber)
	resp, err := c.HTTPClient.Get(reqURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request to user API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var userInfo models.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, errors.Wrap(err, "failed to decode response from user API")
	}

	return &userInfo, nil
}

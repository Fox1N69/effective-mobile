package clients

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

// FetchUserInfo retrieves user information from the API based on the given passport series and number
//
// # If your api returns data in an array
//
// then uncomment the commented code and delete "var userInfos models.UserInfo"
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

	//var userInfos []models.UserInfo
	var userInfos models.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfos); err != nil {
		return nil, errors.Wrap(err, "failed to decode response from user API")
	}

	/*
		if len(userInfos) == 0 {
			return nil, errors.New("no user info found")
		}

		return &userInfos[0], nil
	*/
	return &userInfos, nil
}

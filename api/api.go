package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type uploadthingConfig struct {
	Host    string `json:"host"`
	ApiKey  string
	Version string `json:"version"`
}

type UtApi struct {
	config     *uploadthingConfig
	httpClient *http.Client
}

// NewUtApi - Create a new UtApi instance with the given API key.
//
// Example:
// utapi := utapigo.NewUtApi("sk_...")
func NewUtApi(ApiKey string) *UtApi {
	config := &uploadthingConfig{Host: "https://api.uploadthing.com", ApiKey: ApiKey, Version: "v6"}

	return &UtApi{config: config, httpClient: &http.Client{}}
}

// MakeRequest - Make a request to the uploadthing API.
// Returns the response and an error if any occurred.
// The error will be nil if the request was successful.
// If the request fails, the error will contain an error message.
// The response body will be read and closed automatically.
// The response will contain the HTTP status code.
// If the status code is not 200, the error message will include the response body.
func (u *UtApi) MakeRequest(apiEndpoint string, method string, body *bytes.Buffer) (*http.Response, error) {
	url := u.getUploadthingUrl() + apiEndpoint
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("uploadthing request failed, reason: %v", err)
	}
	req.Header.Add("x-uploadthing-api-key", u.config.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := u.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("uploadthing request failed, reason: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := io.Copy(respBody, resp.Body)
		if err != nil {
			return nil, fmt.Errorf("uploadthing request failed, status code: %d, but couldn't read body: %v", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("uploadthing request failed, status code: %d, body: %s", resp.StatusCode, respBody.String())
	}
	return resp, nil

}

// getUploadthingUrl - Get the full URL for the uploadthing API.
// Returns the full URL for the uploadthing API.
// The URL includes the host, version, and API key.
// If the version is not provided, it defaults to "v6".
// If the host is not provided, it defaults to "https://api.uploadthing.com".
// The API key is included in the request header.
// The URL is constructed using the provided host and version.
func (u *UtApi) getUploadthingUrl() string {
	return u.config.Host + "/" + u.config.Version
}

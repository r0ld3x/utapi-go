package utapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type GetUsageInfoResponse struct {
	TotalBytes    int64 `json:"totalBytes"`
	AppTotalBytes int64 `json:"appTotalBytes"`
	FilesUploaded int64 `json:"filesUploaded"`
	LimitBytes    int64 `json:"limitBytes"`
}

// GetUsageInfo retrieves the total usage information for the user.
// Returns a pointer to a GetUsageInfoResponse object on success, or an error on failure.
//
// Example:
//
//	usageInfo, err := utapi.GetUsageInfo()
//	if err != nil {
//	    fmt.Println("Error fetching usage info:", err.Error())
//	    return
//	}
func (u *UtApi) GetUsageInfo() (*GetUsageInfoResponse, error) {

	apiEndpoint := "/getUsageInfo"
	resp, err := u.MakeRequest(apiEndpoint, "POST", &bytes.Buffer{})
	if err != nil {
		println("Error making request:", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var parsed GetUsageInfoResponse
	err = json.Unmarshal(bodyBytes, &parsed)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &parsed, nil
}

package utapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type ListFilesOpts struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ListFilesResponse struct {
	HasMore bool   `json:"hasMore"`
	Files   []File `json:"files"`
}

type File struct {
	ID         string `json:"id"`
	Key        string `json:"key"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	UploadedAt int64  `json:"uploadedAt"`
	CustomID   any    `json:"customId"`
	Status     string `json:"status"`
}

// ListFiles lists files in the user's storage.
// Example: -
//
//	listFiles, err := client.ListFiles(ut.ListFilesOpts{Limit: 10, Offset: 0})
//	if err != nil {
//	  fmt.Println("Error listing files:", err.Error())
//	  return
//	}
//	fmt.Printf("List of files (Limit: %d, Offset: %d):\n", listFiles.Limit, listFiles.Offset)
func (u *UtApi) ListFiles(opts ListFilesOpts) (*ListFilesResponse, error) {
	payload, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	apiEndpoint := "/listFiles"
	body := bytes.NewBuffer(payload)
	resp, err := u.MakeRequest(apiEndpoint, "POST", body)
	if err != nil {
		println("Error making request:", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var parsed ListFilesResponse
	err = json.Unmarshal(bodyBytes, &parsed)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &parsed, nil
}

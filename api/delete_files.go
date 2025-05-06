package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type DeleteFilesOpt struct {
	FileKeys []string `json:"fileKeys"`
}

type DeleteFilesResponse struct {
	Success      bool  `json:"success"`
	DeletedCount int64 `json:"deletedCount"`
}

// DeleteFiles deletes multiple files by their keys.
// Returns the number of files deleted.
//
// Example:
//
//	deleteResponse, err := client.DeleteFiles(utapi.DeleteFilesOpt{
//	  FileKeys: []string{"file1", "file2", "file3"},
//	})
//	fmt.Println(deleteResponse.DeletedCount) // Output: 3
func (u *UtApi) DeleteFiles(opts DeleteFilesOpt) (*DeleteFilesResponse, error) {
	payload, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	apiEndpoint := "/deleteFiles"
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

	var parsed DeleteFilesResponse
	err = json.Unmarshal(bodyBytes, &parsed)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &parsed, nil
}

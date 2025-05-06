package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type RenameFilesOpts []RenameFilesOpt

type RenameFilesOpt struct {
	NewName string `json:"newName"`
	FileKey string `json:"fileKey"`
}

type RenameRequest struct {
	Updates RenameFilesOpts `json:"updates"`
}

type RenameFilesResponse struct {
	Success      bool  `json:"success"`
	RenamedCount int64 `json:"renamedCount"`
}

// RenameFiles renames files in a user's cloud storage.
// Returns a RenameFilesResponse containing a boolean indicating success and the count of renamed files.
//
// Example:
//
//	 opts := []RenameFilesOpt{
//	   {NewName: "new_name.txt", FileKey: "old_key.txt"},
//	 }
//	 result, err := client.RenameFiles(opts)
//	 if err != nil {
//	   panic(err)
//	}
func (u *UtApi) RenameFiles(opts RenameFilesOpts) (*RenameFilesResponse, error) {
	payload, err := json.Marshal(RenameRequest{Updates: opts})
	if err != nil {
		return nil, err
	}
	apiEndpoint := "/renameFiles"
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

	var parsed RenameFilesResponse
	err = json.Unmarshal(bodyBytes, &parsed)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &parsed, nil
}

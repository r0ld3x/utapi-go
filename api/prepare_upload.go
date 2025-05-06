package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FileRequest struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Type string `json:"type"`
}

type PrepareUploadOpt struct {
	Files        []FileRequest `json:"files"`
	CallbackURL  string        `json:"callbackUrl"`
	CallbackSlug string        `json:"callbackSlug"`
	RouteConfig  []string      `json:"routeConfig"`
}

type PrepareUploadResponse struct {
	URL                string `json:"url"`
	Fields             Fields `json:"fields"`
	Key                string `json:"key"`
	ContentDisposition string `json:"contentDisposition"`
	FileURL            string `json:"fileUrl"`
	AppURL             string `json:"appUrl"`
	UfsURL             string `json:"ufsUrl"`
	PollingJwt         string `json:"pollingJwt"`
	PollingURL         string `json:"pollingUrl"`
	FileName           string `json:"fileName"`
	FileType           string `json:"fileType"`
	CustomID           any    `json:"customId"`
}

type Fields struct {
	ContentType        string `json:"Content-Type"`
	ContentDisposition string `json:"Content-Disposition"`
	Bucket             string `json:"bucket"`
	XAmzAlgorithm      string `json:"X-Amz-Algorithm"`
	XAmzCredential     string `json:"X-Amz-Credential"`
	XAmzDate           string `json:"X-Amz-Date"`
	XAmzSecurityToken  string `json:"X-Amz-Security-Token"`
	Key                string `json:"key"`
	Policy             string `json:"Policy"`
	XAmzSignature      string `json:"X-Amz-Signature"`
}

// PrepareUpload retrieves the data to upload file
// Returns the URL to upload the file, fields for the POST request, and the key for the file
// If there's an error, it returns an error message.
//
// Example:
//
//	prepareUpload, err := client.PrepareUpload(utapi.PrepareUploadOpt{
//	  Files: []utapi.FileRequest{{Name: "file.txt", Size: 1024, Type: "text/plain"}},
//	  CallbackURL:  "http://example.com/callback",
//	  CallbackSlug: "file.txt",
//	  RouteConfig:  []string{"text"},
//	})
//	if err != nil {
//	  fmt.Println("Error preparing upload:", err.Error())
//	  return
//	}
func (u *UtApi) PrepareUpload(opts PrepareUploadOpt) (*PrepareUploadResponse, error) {
	payload, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	apiEndpoint := "/prepareUpload"
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

	var parsed []PrepareUploadResponse
	err = json.Unmarshal(bodyBytes, &parsed)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	return &parsed[0], nil

}

func GetFileInfo(filePath string) (*FileRequest, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}
	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	contentType := http.DetectContentType(buffer[:n])

	return &FileRequest{
		Name: file.Name(),
		Size: fileStat.Size(),
		Type: contentType,
	}, nil
}

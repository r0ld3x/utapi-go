package utapi

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadFile uploads a file to the UploadThing server.
//
// This function takes a file path from the local filesystem and uploads the file
// to the server using the information provided in the PrepareUploadResponse.
//
// Example:
//
//	resp, _ := ut.PrepareUpload(...)
//	err := ut.UploadFile(&resp, "example.pdf")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Parameters:
//   - responseItem: A pointer to the PrepareUploadResponse object returned from the PrepareUpload API.
//   - filePath: The full path to the local file you want to upload.
//
// Returns an error if the upload fails or the file cannot be read.
func (u *UtApi) UploadFile(responseItem *PrepareUploadResponse, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	for key, value := range map[string]string{
		"Content-Type":         responseItem.Fields.ContentType,
		"Content-Disposition":  responseItem.Fields.ContentDisposition,
		"bucket":               responseItem.Fields.Bucket,
		"X-Amz-Algorithm":      responseItem.Fields.XAmzAlgorithm,
		"X-Amz-Credential":     responseItem.Fields.XAmzCredential,
		"X-Amz-Date":           responseItem.Fields.XAmzDate,
		"X-Amz-Security-Token": responseItem.Fields.XAmzSecurityToken,
		"key":                  responseItem.Fields.Key,
		"Policy":               responseItem.Fields.Policy,
		"X-Amz-Signature":      responseItem.Fields.XAmzSignature,
	} {
		if value != "" {
			_ = writer.WriteField(key, value)
		}
	}
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", responseItem.URL, requestBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := u.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error uploading file: %s, status code: %d", string(bodyBytes), resp.StatusCode)
	}
	return nil
}

package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func (c *FileStoreClient) AddFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	headers := map[string]string{
		"X-Filename":     filepath.Base(filePath),
		"Content-Type":   "application/octet-stream",
	}

	resp, err := c.makeRequest("POST", c.BaseURL+"/files", bytes.NewReader(content), headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	if !response["success"].(bool) {
		return fmt.Errorf(response["message"].(string))
	}

	fmt.Printf("Added file: %s\n", filepath.Base(filePath))
	return nil
}

// Implement other methods similarly: ListFiles, RemoveFile, UpdateFile, WordCount, FrequentWords
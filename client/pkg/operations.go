package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	if !response["success"].(bool) {
		return fmt.Errorf(response["message"].(string))
	}

	fmt.Printf("Added file: %s\n", filepath.Base(filePath))
	return nil
}

func (c *FileStoreClient) ListFiles() error {
	resp, err := http.Get(c.BaseURL + "/files")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var files []string
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return err
	}

	fmt.Println("Files in store:")
	for _, file := range files {
		fmt.Println(file)
	}
	return nil
}

func (c *FileStoreClient) RemoveFile(filename string) error {
	req, err := http.NewRequest("DELETE", c.BaseURL+"/files/"+filename, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	if !response["success"].(bool) {
		return fmt.Errorf(response["message"].(string))
	}

	fmt.Printf("Removed file: %s\n", filename)
	return nil
}

func (c *FileStoreClient) UpdateFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.BaseURL+"/files/"+filepath.Base(filePath), bytes.NewReader(content))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	if !response["success"].(bool) {
		return fmt.Errorf(response["message"].(string))
	}

	fmt.Printf("Updated file: %s\n", filepath.Base(filePath))
	return nil
}

func (c *FileStoreClient) WordCount() error {
	resp, err := http.Get(c.BaseURL + "/files/wordcount")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	fmt.Printf("Total word count: %d\n", result["wordCount"])
	return nil
}

func (c *FileStoreClient) FrequentWords(limit int, order string) error {
	url := fmt.Sprintf("%s/files/frequent-words?limit=%d&order=%s", c.BaseURL, limit, order)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var words [][]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&words); err != nil {
		return err
	}

	fmt.Println("Most frequent words:")
	for _, word := range words {
		fmt.Printf("%s: %d\n", word[0], int(word[1].(float64)))
	}
	return nil
}
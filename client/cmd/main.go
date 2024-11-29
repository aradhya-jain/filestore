package main

import (
	"fmt"
	"os"
	"strconv"

	"your_module_path/pkg"
)

func main() {
	client := &pkg.FileStoreClient{BaseURL: "http://localhost:3000"}

	if len(os.Args) < 2 {
		fmt.Println("Usage: store <command> [args]")
		return
	}

	switch os.Args[1] {
	case "add":
		for _, file := range os.Args[2:] {
			if err := client.AddFile(file); err != nil {
				fmt.Printf("Error adding file %s: %v\n", file, err)
			}
		}
	case "ls":
		if err := client.ListFiles(); err != nil {
			fmt.Println("Error listing files:", err)
		}
	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a file to remove")
			return
		}
		if err := client.RemoveFile(os.Args[2]); err != nil {
			fmt.Println("Error removing file:", err)
		}
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a file to update")
			return
		}
		if err := client.UpdateFile(os.Args[2]); err != nil {
			fmt.Println("Error updating file:", err)
		}
	case "wc":
		if err := client.WordCount(); err != nil {
			fmt.Println("Error getting word count:", err)
		}
	case "freq-words":
		limit := 10
		order := "dsc"
		if len(os.Args) > 2 {
			limit, _ = strconv.Atoi(os.Args[2])
		}
		if err := client.FrequentWords(limit, order); err != nil {
			fmt.Println("Error getting frequent words:", err)
		}
	default:
		fmt.Println("Unknown command")
	}
}
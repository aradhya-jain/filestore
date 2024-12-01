package main

import (
	"fmt"
	"os"
	"strconv"
	"filestore/pkg"
)

func main() {

	client := &pkg.FileStoreClient{BaseURL: "http://localhost:3000"}

	if len(os.Args) < 2 {
		fmt.Println("Usage: store <command> [args]")
		fmt.Println("Available Commands")
		fmt.Println("add: Adds a file to the server")
		fmt.Println("ls: Lists all the files in the server")
		fmt.Println("rm: Removes file from the server")
		fmt.Println("update: updates the contents of the file")
		fmt.Println("wc: returns the word count")
		fmt.Println("freq-words: returns the frequency of the words")
		return
	}

	switch os.Args[1] {
	case "add":
		for _, file := range os.Args[2:] {
			if err := client.AddFile(file); err != nil {
				fmt.Printf("Error adding file %s: %v\n", file, err)
				os.Exit(1)
				return
			}
		}
	case "ls":
		if err := client.ListFiles(); err != nil {
			fmt.Println("Error listing files:", err)
			os.Exit(1)
			return
		}
	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a file to remove")
			os.Exit(1)
			return
		}
		if err := client.RemoveFile(os.Args[2]); err != nil {
			fmt.Println("Error removing file:", err)
		}
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a file to update")
			os.Exit(1)
			return
		}
		if err := client.UpdateFile(os.Args[2]); err != nil {
			fmt.Println("Error updating file:", err)
			os.Exit(1)
			return
		}
	case "wc":
		if err := client.WordCount(); err != nil {
			fmt.Println("Error getting word count:", err)
			os.Exit(1)
			return
		}
	case "freq-words":
		limit := 10
		var err error 
		order := "dsc"
		if len(os.Args) > 2 {
			limit, err = strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("error parsing the limit:", err)
				os.Exit(1)
				return
			}
		}
		if err := client.FrequentWords(limit, order); err != nil {
			fmt.Println("Error getting frequent words:", err)
			os.Exit(1)
			return
		}
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
		return
	}
}
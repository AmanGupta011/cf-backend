package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// func MakeFile(path string, fileName string) error {
// 	fmt.Println("hello from Makefile")
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(wd, path)
// 	pathToCreateFile := filepath.Join(wd,path)
// 	fileToCreate := filepath.Join(pathToCreateFile, fileName)
// 	file, err := os.Create(fileToCreate)
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("errror")
// 		fmt.Println(fileToCreate)
// 		return err
// 	}
// 	file.Close()
// 	return nil
// }

func MakeFile(path string, fileName string) error {
    // Print debug message
    fmt.Println("Creating file:", fileName, "in directory:", path)

    // Get current working directory
    wd, err := os.Getwd()
    if err != nil {
        return err
    }

    // Construct absolute path to create file
    absPath := filepath.Join(wd, path)
    fullPath := filepath.Join(absPath, fileName)

    // Ensure the directory exists, create it if necessary
    if err := os.MkdirAll(absPath, 0755); err != nil {
        return fmt.Errorf("failed to create directory: %s", err)
    }

    // Create the file
    file, err := os.Create(fullPath)
    if err != nil {
        return fmt.Errorf("failed to create file: %s", err)
    }
    defer file.Close()

    // Print success message
    fmt.Println("File created:", fullPath)

    return nil
}

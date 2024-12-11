package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: code-map <path> <pattern regex> <optional output file>")
		return
	}

	rootPath := os.Args[1]
	pattern := os.Args[2]
	outputFile := "project.md"

	if len(os.Args) > 3 {
		outputFile = os.Args[3]
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("invalid regex pattern: %v\n", err)
		return
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("failed to create output file: %v\n", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error accessing path %s: %v\n", path, err)
			return nil
		}

		if !info.IsDir() && regex.MatchString(info.Name()) {
			relativePath, _ := filepath.Rel(rootPath, path)
			if err = processFile(path, relativePath, writer); err != nil {
				fmt.Printf("error processing file %s: %v\n", path, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking through the directory: %v\n", err)
		return
	}

	_ = writer.Flush()
	fmt.Printf("project code-map saved to %s\n", outputFile)
}

func processFile(filePath, relativePath string, writer *bufio.Writer) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	header := fmt.Sprintf("###%s###\n```%s\n", relativePath, getFileExtension(filePath))
	_, _ = writer.WriteString(header)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, _ = writer.WriteString(scanner.Text() + "\n")
	}

	if err = scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	_, _ = writer.WriteString("```\n\n")
	return nil
}

func getFileExtension(filePath string) string {
	parts := strings.Split(filePath, ".")
	if len(parts) < 2 {
		return ""
	}
	return parts[len(parts)-1]
}

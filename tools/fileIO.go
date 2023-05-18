package tools

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// WriteToFile writes the given string of data to the specified file path
func WriteToFile(filePath string, data string) error {
	path := GetAbsolutePathFrom(filePath)
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

// GetAbsolutePathFrom returns the absolute path from a relative path
func GetAbsolutePathFrom(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, path)
}

// CreateDirectory creates a directory at the specified path
func CreateDirectory(filePath string) error {
	path := GetAbsolutePathFrom(filePath)

	if FileExist(path) {
		return fmt.Errorf("directory already exists: %s", filePath)
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// FileExist checks if a file exists at the specified path
func FileExist(filePath string) bool {
	path := GetAbsolutePathFrom(filePath)
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

// ReadFile reads the file at filePath and returns its contents as a byte slice.
func ReadFile(filePath string) ([]byte, error) {
	path := GetAbsolutePathFrom(filePath)

	if !FileExist(path) {
		return nil, fmt.Errorf("file does not exist: %s", filePath)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %s: %w", filePath, err)
	}

	return data, nil
}

// GetDirectoriesFrom returns a list of directories from a file path
func GetDirectoriesFrom(filePath string) ([]string, error) {
	path := GetAbsolutePathFrom(filePath)
	if !FileExist(path) {
		return nil, fmt.Errorf("file does not exist: %s", filePath)
	}

	directory, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %s: %w", filePath, err)
	}

	files := make([]string, 0, len(directory))
	for _, file := range directory {
		if file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

// GetFilesFrom returns a list of files from a file path
func GetFilesFrom(filePath string) ([]string, error) {
	path := GetAbsolutePathFrom(filePath)
	if !FileExist(path) {
		return nil, fmt.Errorf("file does not exist: %s", filePath)
	}

	directory, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %s: %w", filePath, err)
	}

	files := make([]string, 0, len(directory))
	for _, file := range directory {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

func TransformInterfaceIntoMappedArray(data []interface{}) []map[string]interface{} {
	results := make([]map[string]interface{}, 0, len(data))
	for _, v := range data {
		result := v.(map[string]interface{})
		results = append(results, result)
	}
	return results
}

func TransformInterfaceIntoMappedObject(data interface{}) map[string]interface{} {
	result := data.(map[string]interface{})
	return result
}

func AuditArrayCapacity(data []map[string]interface{}) []map[string]interface{} {
	dataLen := len(data)
	results := make([]map[string]interface{}, 0, dataLen)
	for i := 0; i < dataLen; i++ {
		result := data[i]
		results = append(results, result)
	}
	return results
}

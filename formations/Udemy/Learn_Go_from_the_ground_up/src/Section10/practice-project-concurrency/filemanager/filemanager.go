package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	inputFilePath  string
	outputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.inputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open File")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.New("Failed to read line in File")
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	createdFile, err := os.Create(fm.outputFilePath)

	if err != nil {
		return errors.New("Failed to create file!")
	}

	defer createdFile.Close()

	time.Sleep(3 * time.Second)

	encoderValue := json.NewEncoder(createdFile)
	err = encoderValue.Encode(data)

	if err != nil {
		return errors.New("Failed to encode data!")
	}

	return nil
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		inputFilePath:  inputFilePath,
		outputFilePath: outputFilePath,
	}
}

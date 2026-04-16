package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
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

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read line in File")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	createdFile, err := os.Create(fm.outputFilePath)

	if err != nil {
		return errors.New("Failed to create file!")
	}

	encoderValue := json.NewEncoder(createdFile)
	err = encoderValue.Encode(data)

	if err != nil {
		createdFile.Close()
		return errors.New("Failed to encode data!")
	}

	createdFile.Close()
	return nil
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		inputFilePath:  inputFilePath,
		outputFilePath: outputFilePath,
	}
}

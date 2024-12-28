package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct{
	InputPath string
	OutputPath string
}
func New(inputPath, outputPath string) FileManager{
	return FileManager{
		InputPath: inputPath,
		OutputPath: outputPath,
	}
}

func (fm FileManager) ReadData() ([]string, error) {
	file, err := os.Open(fm.InputPath)

	if(err != nil){
		return nil, errors.New("Failed to open the file...!")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	scanErr := scanner.Err()
	if(scanErr != nil){
		file.Close()
		return nil, errors.New("Failed to read the file...!")
	}
	file.Close()
  return lines, nil
}

func (fm FileManager) WriteJsonData(data any) error{
	file, err := os.Create(fm.OutputPath)
	if(err != nil){
		return errors.New("Failed to create the file...!")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if(err != nil){
		file.Close()
		return errors.New("Failed to contvert data to JSON...!")
	}
  return nil
}
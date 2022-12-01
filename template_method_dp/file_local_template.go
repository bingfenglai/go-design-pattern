package template_method_dp

import (
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type LocalFileTemplate struct {
}

func (l *LocalFileTemplate) LoadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)

	defer file.Close()

	if err != nil {
		return "", err
	}

	return file.Name(), nil

}

func (l *LocalFileTemplate) SaveFile(filePath string) (string, error) {

	file, err := os.Open(filePath)

	defer file.Close()

	if err != nil {
		return "", err
	}

	fileOriginalName := file.Name()

	index := strings.LastIndex(fileOriginalName, ".")

	fileExtension := fileOriginalName[index:]

	nano := time.Now().UnixNano()
	writeFileName := strconv.FormatInt(nano, 16) + fileExtension
	targetFile, err := os.Create(writeFileName)

	if err != nil {
		return "", err
	}

	_, err = io.Copy(targetFile, file)
	if err != nil {
		return "", err
	}

	return writeFileName, nil
}

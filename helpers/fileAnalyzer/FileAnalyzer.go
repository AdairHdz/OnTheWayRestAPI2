package fileAnalyzer

import (
	"io"
	"os")


func DirIsEmpty(fileName string) (bool, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return false, err
	}

	defer file.Close()

	_, err = file.Readdir(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err
}

func ImageHasValidFormat(fileExtension string) bool {
	if fileExtension == ".png" || fileExtension == ".jpeg" || fileExtension == ".jpg" {
		return true
	}

	return false
}

func EvidenceHasValidFormat(fileExtension string) bool {
	isImageWithValidFormat := ImageHasValidFormat(fileExtension)
	isVideoWithValidFormat := false

	if fileExtension == ".mp4" {
		return true
	}

	if !isImageWithValidFormat && !isVideoWithValidFormat {
		return false
	}
	return true
}
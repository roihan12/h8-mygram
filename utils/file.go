package utils

import (
	"mime/multipart"
	"strings"
)

func CheckFile(file *multipart.FileHeader) error {
	_, err := file.Open()
	if err != nil {
		return ErrFormatFile
	}

	err = CheckFileSize(file.Size)
	if err != nil {

		return ErrFileSize
	}

	_, err = CheckFileExtensionImage(file.Filename)
	if err != nil {
		return ErrFileExtension
	}

	return nil
}

func CheckFileExtensionImage(filename string) (string, error) {
	extension := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])

	if extension != "jpg" && extension != "jpeg" && extension != "png" {
		return "", ErrFileExtension
	}
	return extension, nil
}

func CheckFileSize(size int64) error {
	var fileSize int64 = 1097152
	if size == 0 {
		return ErrFileSize
	}

	if size > fileSize {
		return ErrFileSize
	}

	return nil
}

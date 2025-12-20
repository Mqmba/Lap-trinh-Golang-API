package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var allowedExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

var allowedMimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

const maxSize = 5 << 20

func ValidateAndSaveFile(fileHeader *multipart.FileHeader, uploadDir string) (string, error) {

	// Check extension in filename
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExts[ext] {
		return "", errors.New("Unsupported file extension")
	}

	// Check size
	if fileHeader.Size > maxSize {
		return "", errors.New("File is too large (max 5 MB)")
	}

	// Check file type
	file, err := fileHeader.Open()
	if err != nil {
		return "", errors.New("Cannot open file")
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", errors.New("Cannot read file")
	}

	mimeType := http.DetectContentType(buffer)
	if !allowedMimeTypes[mimeType] {
		return "", fmt.Errorf("Invalid MIME type : %s", mimeType)
	}

	// Change file name
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// Create folder if not exist
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		return "", errors.New("Cannot create upload folder")
	}

	// uploadDir "./upload" + filename "abc.jpg"
	savePath := filepath.Join(uploadDir, filename)
	if err := saveFile(fileHeader, savePath); err != nil {
		return "", err
	}

	return filename, nil
}

func saveFile(fileHeader *multipart.FileHeader, destination string) error {
	// Mở file hiện tại
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Tạo file đích đến
	// Di chuyển nội dung file vào file đích đến
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	return err
}

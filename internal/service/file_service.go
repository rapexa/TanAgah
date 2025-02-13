package service

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"TanAgah/internal/config"
	"TanAgah/internal/repository"
)

type FileService struct {
	repo   repository.FileRepository
	config *config.Config
}

func NewFileService(repo repository.FileRepository, cfg *config.Config) *FileService {
	return &FileService{
		repo:   repo,
		config: cfg,
	}
}

func (s *FileService) ProcessUpload(fileHeader *multipart.FileHeader, userId string) (string, error) {
	// Validate file size
	if fileHeader.Size > s.config.MaxUploadSize {
		return "", fmt.Errorf("file size exceeds limit")
	}

	// Validate file type
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file")
	}
	defer file.Close()

	// Validate MIME type
	buff := make([]byte, 512)
	if _, err = file.Read(buff); err != nil {
		return "", fmt.Errorf("invalid file content")
	}
	if !s.isValidMIMEType(buff) {
		return "", fmt.Errorf("unsupported file type")
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), filepath.Base(fileHeader.Filename))

	// Reset file reader
	if _, err = file.Seek(0, 0); err != nil {
		return "", fmt.Errorf("failed to read file")
	}

	// Save file
	if err := s.repo.SaveFile(filename, file, userId); err != nil {
		return "", fmt.Errorf("failed to save file")
	}

	return filename, nil
}

func (s *FileService) isValidMIMEType(buff []byte) bool {
	mimeType := http.DetectContentType(buff)
	for _, allowed := range s.config.AllowedTypes {
		if mimeType == allowed {
			return true
		}
	}
	return false
}

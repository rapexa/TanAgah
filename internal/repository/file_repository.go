package repository

import (
	"io"
	"os"
	"path/filepath"
)

type FileRepository interface {
	SaveFile(filename string, content io.Reader, userId string) error
	GenerateUniqueFilename(originalName string) string
}

type LocalFileRepository struct {
	uploadDir string
}

func NewFileRepository(uploadDir string) *LocalFileRepository {
	return &LocalFileRepository{uploadDir: uploadDir}
}

func (r *LocalFileRepository) SaveFile(filename string, content io.Reader, userId string) error {
	dstPath := filepath.Join(r.uploadDir, filename+"-"+userId)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, content)
	return err
}

func (r *LocalFileRepository) GenerateUniqueFilename(originalName string) string {
	return filepath.Base(originalName)
}

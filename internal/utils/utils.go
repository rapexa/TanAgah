package utils

import (
	"path/filepath"
	"strings"
)

// Helper function
func IsValidFileType(filename string) bool {
	validExtensions := []string{".jpg", ".png", ".jpeg"}
	ext := strings.ToLower(filepath.Ext(filename))
	for _, validExt := range validExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}

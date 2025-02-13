package model

type UploadResponse struct {
	Status  string   `json:"status"`
	Files   []string `json:"files"`
	Message string   `json:"message,omitempty"`
}

type FileInfo struct {
	Name        string
	Size        int64
	ContentType string
}

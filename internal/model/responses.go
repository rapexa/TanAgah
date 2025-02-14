package model

type MainRp struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UploadResponse struct {
	Status  string   `json:"status"`
	Files   []string `json:"files"`
	Message string   `json:"message,omitempty"`
}

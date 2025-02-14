package model

type MainRp struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	JwtToken string `json:"jwt_token"`
}

type UploadResponse struct {
	Status  string   `json:"status"`
	Files   []string `json:"files"`
	Message string   `json:"message,omitempty"`
}

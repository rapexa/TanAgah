package controller

import (
	"TanAgah/internal/model"
	"TanAgah/internal/service"
	"TanAgah/internal/stringResource"
	"TanAgah/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileController struct {
	fileService service.FileService
	userService service.UserService // Add UserService dependency
}

func NewFileController(fs service.FileService, us service.UserService) *FileController {
	return &FileController{
		fileService: fs,
		userService: us,
	}
}

func (cf *FileController) HandleFileUpload(ctx *gin.Context) {

	id := ctx.Param("id")

	// First get the user
	user, err := cf.userService.GetUser(id)

	if err != nil {
		utils.SendError404Response(ctx, stringResource.GetStrings().UserNotFound(ctx))
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.UploadResponse{
			Status:  "error",
			Message: "Invalid request format",
		})
		return
	}

	files := form.File["file"]
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, model.UploadResponse{
			Status:  "error",
			Message: "No files uploaded",
		})
		return
	}

	var uploadedFiles []string
	var errors []string

	for _, fileHeader := range files {
		filename, err := cf.fileService.ProcessUpload(fileHeader, user.Name)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: %v", fileHeader.Filename, err))
			continue
		}
		uploadedFiles = append(uploadedFiles, filename)
	}

	if len(errors) > 0 {
		ctx.JSON(http.StatusPartialContent, model.UploadResponse{
			Status:  "partial",
			Files:   uploadedFiles,
			Message: fmt.Sprintf("%d files failed to upload", len(errors)),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.UploadResponse{
		Status: "success",
		Files:  uploadedFiles,
	})
}

package controller

import (
	"github.com/gin-gonic/gin"
)

// FileController manages the file resource
type FileController interface {
	ReadImageHandler(*gin.Context)
	CreateImageHandler(*gin.Context)
}
type fileController struct {
}

// File returns file resolver
func (c *Controllers) File() FileController {
	return &fileController{}
}

func (c *fileController) ReadImageHandler(cxt *gin.Context) {
}

func (c *fileController) CreateImageHandler(cxt *gin.Context) {
}

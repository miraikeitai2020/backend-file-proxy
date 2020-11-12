package controller

import "github.com/gin-gonic/gin"

// MinioController manages the minio resource
type MinioController interface {
	ConfigUpdateHandler(*gin.Context)
}
type minioController struct {
}

// Minio returns minio resolver
func (c *Controllers) Minio() MinioController {
	return &minioController{}
}

func (c *minioController) ConfigUpdateHandler(cxt *gin.Context) {

}

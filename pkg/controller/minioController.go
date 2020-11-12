package controller

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/dao/minio"
	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/dto"
	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/service/log"
	"github.com/miraikeitai2020/backend-file-proxy/pkg/view"
)

// FileController manages the file resource
type FileController interface {
	ReadImageHandler(*gin.Context)
	CreateImageHandler(*gin.Context)
	ConfigUpdateHandler(*gin.Context)
}
type fileController struct {
	Minio *minio.Minio
}

// Minio returns file resolver
func (c *Controllers) Minio() FileController {
	return &fileController{Minio: c.minio}
}

func (c *fileController) ReadImageHandler(cxt *gin.Context) {
	fileName := cxt.Param("id") + ".jpg"

	object, size, err := c.Minio.Buckets.Detour.Get(fileName)
	if err != nil {
		log.Error(err)
		return
	}
	defer object.Close()

	log.Info(fmt.Sprintf("Create local file (%s)", fileName))
	local, err := os.Create(fileName)
	if err != nil {
		log.Error(err)
		return
	}
	defer local.Close()

	log.Info(fmt.Sprintf("Copy data to local file (%s)", fileName))
	if _, err := io.CopyN(local, object, size); err != nil {
		log.Error(err)
		return
	}

	cxt.File(fileName)
	log.Info(fmt.Sprintf("Remove local file (%s)", fileName))
	if err := os.Remove(fileName); err != nil {
		log.Error(err)
		return
	}
}

func (c *fileController) CreateImageHandler(cxt *gin.Context) {
}

func (c *fileController) ConfigUpdateHandler(cxt *gin.Context) {
	var request dto.ConfigUpdateRequest
	if err := cxt.BindJSON(&request); err != nil {
		log.Error(err)
		appErr := view.NewAppError(dto.ERROR_CODE_INTERNAL, err)
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}

	// check response value
	if request.URL == "" {
		appErr := view.NewAppError(dto.ERROR_CODE_CLIENT, errors.New("`url` parameter value is empty"))
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}
	if request.PublicKey == "" {
		appErr := view.NewAppError(dto.ERROR_CODE_CLIENT, errors.New("`publicKey` parameter value is empty"))
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}
	if request.SecretKey == "" {
		appErr := view.NewAppError(dto.ERROR_CODE_CLIENT, errors.New("`secretKey` parameter value is empty"))
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}

	c.Minio.UpdateConfig(request.URL, request.PublicKey, request.SecretKey)

	cxt.JSON(http.StatusOK, gin.H{"minioConfig": c.Minio.Config})
}

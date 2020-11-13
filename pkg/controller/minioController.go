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
	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/service/decode"
	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/service/log"
	"github.com/miraikeitai2020/backend-file-proxy/pkg/view"
)

// MinioController manages the file resource
type MinioController interface {
	ReadDetourImageHandler(*gin.Context)
	CreateDetourImageHandler(*gin.Context)
	ConfigUpdateHandler(*gin.Context)
	InitMinioHandler(*gin.Context)
}
type minioController struct {
	Minio *minio.Minio
}

// Minio returns file resolver
func (c *Controllers) Minio() MinioController {
	return &minioController{Minio: c.minio}
}

func (c *minioController) ReadDetourImageHandler(cxt *gin.Context) {
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

func (c *minioController) CreateDetourImageHandler(cxt *gin.Context) {
	var request dto.CreateImageRequest

	if err := cxt.BindJSON(&request); err != nil {
		log.Error(err)
		appErr := view.NewAppError(dto.ERROR_CODE_INTERNAL, err)
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}

	fileName := request.ID + ".jpg"

	// check request parameter
	if request.ID == "" {
		appErr := view.NewAppError(dto.ERROR_CODE_CLIENT, errors.New("`id` parameter value is empty"))
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}
	if request.Source == "" {
		appErr := view.NewAppError(dto.ERROR_CODE_CLIENT, errors.New("`source` parameter value is empty"))
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}

	if err := decode.Image(fileName, request.Source); err != nil {
		log.Error(err)
		appErr := view.NewAppError(dto.ERROR_CODE_INTERNAL, err)
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}

	size, err := c.Minio.Buckets.Detour.Add(fileName)
	if err != nil {
		log.Error(err)
		appErr := view.NewAppError(dto.ERROR_CODE_INTERNAL, err)
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}

	info := view.NewObjectInfo(request.ID, size)
	cxt.JSON(http.StatusOK, gin.H{"objectInfo": info})
	/*
		if err := os.Remove(fileName); err != nil {
			log.Error(err)
			return
		}
	*/
}

func (c *minioController) ConfigUpdateHandler(cxt *gin.Context) {
	var request dto.ConfigUpdateRequest
	var publicKey, secretKey string

	// check access keys
	if publicKey = cxt.GetHeader("Public-Key"); publicKey != c.Minio.Config.PublicKey {
		err := errors.New("`Public-Key` is an invalid value")
		appErr := view.NewAppError(dto.ERROR_CODE_CLIENT, err)
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}
	if secretKey = cxt.GetHeader("Secret-Key"); secretKey != c.Minio.Config.SecretKey {
		err := errors.New("`Secret-Key` is an invalid value")
		appErr := view.NewAppError(dto.ERROR_CODE_CLIENT, err)
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}

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

func (c *minioController) InitMinioHandler(cxt *gin.Context) {
	var publicKey, secretKey string

	// check access keys
	if publicKey = cxt.GetHeader("Public-Key"); publicKey != c.Minio.Config.PublicKey {
		err := errors.New("`Public-Key` is an invalid value")
		appErr := view.NewAppError(dto.ERROR_CODE_CLIENT, err)
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}
	if secretKey = cxt.GetHeader("Secret-Key"); secretKey != c.Minio.Config.SecretKey {
		err := errors.New("`Secret-Key` is an invalid value")
		appErr := view.NewAppError(dto.ERROR_CODE_CLIENT, err)
		cxt.JSON(http.StatusOK, gin.H{"error": appErr})
		return
	}
	// create new bucket
	for _, bucket := range minio.BucketList {
		if err := c.Minio.CreateBucket(bucket); err != nil {
			appErr := view.NewAppError(dto.ERROR_CODE_INTERNAL, err)
			cxt.JSON(http.StatusOK, gin.H{"error": appErr})
			return
		}
	}

	cxt.JSON(http.StatusOK, gin.H{"bucketList": minio.BucketList})
}

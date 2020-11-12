package controller

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/dao/minio"
	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/service/log"
)

// FileController manages the file resource
type FileController interface {
	ReadImageHandler(*gin.Context)
	CreateImageHandler(*gin.Context)
}
type fileController struct {
	Minio minio.Repositories
}

// File returns file resolver
func (c *Controllers) File() FileController {
	return &fileController{Minio: c.minio.Buckets}
}

func (c *fileController) ReadImageHandler(cxt *gin.Context) {
	fileName := cxt.Param("id") + ".jpg"

	object, size, err := c.Minio.Detour.Get(fileName)
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

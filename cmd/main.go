package main

import (
	"net/http"
	"os"

	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/service/log"

	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/dao/minio"

	"github.com/gin-gonic/gin"
	"github.com/miraikeitai2020/backend-file-proxy/pkg/controller"
)

func router(ctrl controller.Controllers) *gin.Engine {
	router := gin.Default()
	// ping check test
	router.GET("/", func(cxt *gin.Context) {
		cxt.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	// controlle client config resource
	router.PUT("/minio/config/update", ctrl.Minio().ConfigUpdateHandler)
	// controlle file resource
	router.GET("/image/read/:id", ctrl.Minio().ReadImageHandler)
	router.POST("/image/create", ctrl.Minio().CreateImageHandler)
	return router
}

func main() {
	client, err := minio.New()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ctrl := controller.New(client)
	if err := router(ctrl).Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

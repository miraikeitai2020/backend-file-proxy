package main

import (
	"net/http"

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
	router.POST("/minio/config/update", ctrl.Minio().ConfigUpdateHandler)
	// controlle file resource
	router.GET("/image/read/:id", ctrl.File().ReadImageHandler)
	router.POST("/image/create", ctrl.File().CreateImageHandler)
	return router
}

func main() {
	ctrl := controller.New()
	if err := router(ctrl).Run(); err != nil {

	}
}

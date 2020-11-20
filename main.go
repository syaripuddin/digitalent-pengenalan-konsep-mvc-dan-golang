package main

import (
	"github.com/gin-gonic/gin"
	"github.com/syaripuddin/digitalent-pengenalan-konsep-mvc-dan-golang/app/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.POST("/api/v1/antrian", controller.AddAntrianHandler)
	router.GET("/api/v1/antrian/status", controller.GetAntrianHandler)
	router.PUT("api/v1/antrian/id/:idAntrian", controller.UpdateAntrianHandler)
	router.DELETE("api/v1/antrian/id/:idAntrian/delete", controller.DeleteAntrianHandler)
	router.GET("/antrian", controller.PageAntrianHandler)
	router.Run(":8080")
}

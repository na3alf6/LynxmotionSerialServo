package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/na3alf6/LynxmotionSerialServo/src/api/adapters"
)

// Router is storing gin.Engine
var Router *gin.Engine

func initRouter() {
	router := gin.Default()

	controller := adapters.NewServoController()

	router.GET("/Servos/:id", func(c *gin.Context) { ServoController.Show(c) })

	Router = router
}

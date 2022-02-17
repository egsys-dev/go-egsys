package routes

import (
	"github.com/fabiom2211/go-egsys/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/fleets", controllers.ListFleets)
	r.POST("/fleets", controllers.CreateFleets)
	r.GET("/fleets/:id/alerts", controllers.GetFleets)
	r.POST("/fleets/:id/alerts", controllers.CreateFleetAlerts)
	r.GET("/vehicles", controllers.ListVehicles)
	r.POST("/vehicles", controllers.CreateVehicles)
	r.GET("/vehicles/:id/positions", controllers.GetVehiclesPosition)
	r.Run(":8077")
}

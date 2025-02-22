package routes

import (
	"VisualisasiData/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TableRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/api/Table", controllers.INNERJOINdataKecelakaan)

}

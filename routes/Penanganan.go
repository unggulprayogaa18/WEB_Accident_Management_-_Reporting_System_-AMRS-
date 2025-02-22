package routes

import (
	"VisualisasiData/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PenanganRoutes(r *gin.Engine, db *gorm.DB) {
	penangan := r.Group("/api/penangan")
	{
		penangan.GET("/", controllers.GetPengaduan)
		penangan.DELETE("/:id", func(c *gin.Context) { controllers.DeletePenangan(c, db) })
	}

}

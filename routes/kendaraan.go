package routes

import (
	"VisualisasiData/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func KendaraanRoutes(r *gin.Engine, db *gorm.DB) {
	// Enable CORS
	// This will allow all origins, you can customize as needed

	kendaraan := r.Group("/api/kendaraan")
	{
		kendaraan.POST("/", controllers.CreateKendaraan)      // Create a new kendaraan
		kendaraan.GET("/:id", controllers.GetKendaraan)       // Get a kendaraan by ID
		kendaraan.GET("/", controllers.GetAllKendaraan)       // Get all kendaraan records
		kendaraan.PUT("/:id", controllers.UpdateKendaraan)    // Update a kendaraan by ID
		kendaraan.DELETE("/:id", controllers.DeleteKendaraan) // Delete a Lokasi by ID
	}

}

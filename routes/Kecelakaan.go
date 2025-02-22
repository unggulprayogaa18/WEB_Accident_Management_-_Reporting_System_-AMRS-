package routes

import (
	"VisualisasiData/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func KecelakaanRoutes(r *gin.Engine, db *gorm.DB) {
	// Enable CORS
	// This will allow all origins, you can customize as needed

	kecelakaan := r.Group("/api/kecelakaan")
	{
		kecelakaan.POST("/", controllers.CreateKecelakaan)                 // Create a new kecelakaan
		kecelakaan.GET("/:id", controllers.GetKecelakaan)                  // Get a kecelakaan by ID
		kecelakaan.GET("/", controllers.GetAllKecelakaan)                  // Get all kecelakaan records
		kecelakaan.GET("/tahunini", controllers.GetDatakecelakaanTahunini) // Get all kecelakaan records
		kecelakaan.GET("/semuatahun", controllers.Getdatakecelakaansemuatahun)
		kecelakaan.GET("/mobil", controllers.HitungTotalKecelakaanMobileTahunIni)
		kecelakaan.GET("/per-lokasi", controllers.HitungTotalKecelakaanPerLokasi) // Get all kecelakaan records
		// Get all kecelakaan records
		kecelakaan.GET("/penyebab-tertinggi", controllers.HitungPenyebabTertinggi) // Get all kecelakaan records
		kecelakaan.PUT("/:id", controllers.UpdateKecelakaan)                       // Update a kecelakaan by ID
		kecelakaan.DELETE("/:id", controllers.DeleteKecelakaan)                    // Delete a Lokasi by ID
	}

}

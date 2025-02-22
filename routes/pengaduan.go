package routes

import (
	"VisualisasiData/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PengaduanRoutes(r *gin.Engine, db *gorm.DB) {
	// Enable CORS
	// This will allow all origins, you can customize as needed

	Pengaduan := r.Group("/api/Pengaduan")
	{
		Pengaduan.POST("/", controllers.CreatePengaduan)
		Pengaduan.GET("/", controllers.GetAllPengaduan)
		Pengaduan.POST("/:idPengaduan/valid", controllers.UpdateStatusValid)           // Route untuk valid
		Pengaduan.POST("/:idPengaduan/tidakvalid", controllers.UpdateStatusTidakValid) // Route untuk tidak valid
	}

}

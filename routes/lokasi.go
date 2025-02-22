package routes

import (
	"VisualisasiData/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LokasiRoutes(r *gin.Engine, db *gorm.DB) {
	// Enable CORS
	// This will allow all origins, you can customize as needed

	lokasi := r.Group("/api/lokasi")
	{
		lokasi.POST("/", controllers.CreateLokasi)      // Create a new Lokasi
		lokasi.GET("/:id", controllers.GetLokasi)       // Get a Lokasi by ID
		lokasi.GET("/", controllers.GetAllLokasi)       // Get all Lokasi records
		lokasi.PUT("/:id", controllers.UpdateLokasi)    // Update a Lokasi by ID
		lokasi.DELETE("/:id", controllers.DeleteLokasi) // Delete a Lokasi by ID
	}

	// Endpoint to generate Lokasi ID
	r.GET("/api/generate-lokasi-id", func(c *gin.Context) {
		// Menambahkan log untuk memulai proses
		fmt.Println("Received request to generate Lokasi ID")

		// Memanggil fungsi untuk menghasilkan ID baru
		newID := controllers.GenerateLokasiID(db)

		// Debug log untuk memastikan ID baru yang dihasilkan
		fmt.Println("Generated new ID:", newID)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"idLokasi": newID})
		// Pastikan data yang dikirim adalah benar
	})

}

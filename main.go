package main

import (
	"VisualisasiData/database"
	"VisualisasiData/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	db := database.ConnectDB() // Store the DB connection in db

	// Inisialisasi router
	r := gin.Default()
	// Menentukan folder template (folder layout)

	r.Static("/static", "./static")
	r.Static("/frontend", "./frontend")

	// Route untuk file index.html
	r.GET("/", func(c *gin.Context) {
		fmt.Println("Rendering index.html")
		c.File("./frontend/index.html")
	})

	// Register LokasiRoutes with the database connection
	routes.LokasiRoutes(r, db) // Pass the db connection

	// Register KendaraanRoutes with the database connection
	routes.KendaraanRoutes(r, db)

	// Register PengaduanRoutes with the database connection
	routes.PengaduanRoutes(r, db)

	// Register KecelakaanRoutes with the database connection
	routes.KecelakaanRoutes(r, db)

	// Register PenanganRoutes with the database connection
	routes.PenanganRoutes(r, db)

	// Register TableRoutes with the database connection
	routes.TableRoutes(r, db)

	// Register TerkaitdataRoute with the database connection
	routes.TerkaitdataRoute(r)

	// Register UserRoutes with the database connection
	routes.UserRoutes(r)

	// Start server di port 8080
	r.Run(":8080")
}

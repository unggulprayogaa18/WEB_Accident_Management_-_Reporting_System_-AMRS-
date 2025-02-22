package controllers

import (
	"VisualisasiData/database"
	"VisualisasiData/models"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GenerateLokasiID(db *gorm.DB) string {
	var lokasiIDs []string

	// Query untuk mengambil semua ID Lokasi dengan prefix "PB"
	query := `SELECT idLokasi FROM lokasi WHERE deleted_at IS NULL AND idLokasi LIKE '%-PB'`
	err := db.Raw(query).Scan(&lokasiIDs).Error

	// Debug log untuk memastikan query berhasil mengambil semua ID
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return ""
	} else {
		log.Printf("All Lokasi IDs: %+v", lokasiIDs)
	}

	// Tentukan ID terbesar berdasarkan angka tiga digit
	var maxIDNumber int
	for _, id := range lokasiIDs {
		var lastIDNumber int
		_, err := fmt.Sscanf(id, "%d-PB", &lastIDNumber)
		if err != nil {
			log.Printf("Error parsing ID: %v", err)
			continue
		}
		if lastIDNumber > maxIDNumber {
			maxIDNumber = lastIDNumber
		}
	}

	// Generate ID baru dengan menambahkan 1 pada ID terbesar
	newID := fmt.Sprintf("%d-PB", maxIDNumber+1)
	log.Printf("Generated new ID: %s", newID)
	return newID
}

// CreateLokasi handles creating a new Lokasi using raw SQL
func CreateLokasi(c *gin.Context) {
	var lokasi models.Lokasi
	// Parse the JSON body into the lokasi object
	if err := c.ShouldBindJSON(&lokasi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the database.DB instance directly, as it has been initialized in the database package
	db := database.DB

	// Construct the SQL query for inserting the new lokasi record
	query := "INSERT INTO lokasi (idLokasi, namaLokasi, mapLokasi) VALUES (?, ?, ?)"
	result := db.Exec(query, lokasi.IdLokasi, lokasi.NamaLokasi, lokasi.MapLokasi)

	// Check for errors during the execution of the query
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create lokasi", "details": result.Error.Error()})
		return
	}

	// Return a successful response
	c.JSON(http.StatusOK, gin.H{"message": "Lokasi created successfully", "data": lokasi})
}

// GetLokasi retrieves a single Lokasi by ID
func GetLokasi(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Received ID:", id) // Debug: log the received ID

	// Check if the ID is provided in the request
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is missing"})
		return
	}

	var lokasi models.Lokasi
	db := database.DB

	// Define the query to retrieve all columns using the provided 'id'
	query := `
        SELECT idLokasi, namaLokasi, mapLokasi, created_at, updated_at, deleted_at 
        FROM lokasi 
        WHERE idLokasi = ? AND deleted_at IS NULL
        LIMIT 1
    `

	// Log the query for debugging
	fmt.Println("Executing query:", query, " with ID:", id)

	// Execute the raw query and scan the result into 'lokasi'
	if err := db.Raw(query, id).Scan(&lokasi).Error; err != nil {
		fmt.Println("Error executing query:", err) // Debug: log the error if the query fails
		c.JSON(http.StatusNotFound, gin.H{"error": "Lokasi not found"})
		return
	}

	// Log the result for debugging
	fmt.Println("Query result:", lokasi)

	// Return the 'lokasi' struct with all data
	c.JSON(http.StatusOK, gin.H{"data": lokasi})
}

// GetAllLokasi retrieves all Lokasi records with a specific query
// GetAllLokasi retrieves all Lokasi records with a specific query
func GetAllLokasi(c *gin.Context) {
	var lokasi []models.Lokasi
	db := database.DB

	fmt.Println("DEBUG: GetAllLokasi function called")

	// Get the page and limit parameters from the query
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "7")

	// Convert page and limit to integers
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	// Calculate the offset for pagination
	offset := (pageInt - 1) * limitInt

	// Start the query with proper columns (without aliases)
	query := db.Model(&models.Lokasi{}).Limit(limitInt).Offset(offset)

	// Filter by "namaLokasi" query parameter if exists
	namaLokasi := c.Query("namaLokasi")
	if namaLokasi != "" {
		fmt.Printf("DEBUG: Filter applied for namaLokasi: %s\n", namaLokasi)
		query = query.Where("namaLokasi LIKE ?", "%"+namaLokasi+"%")
	}

	// Execute the query
	if err := query.Find(&lokasi).Error; err != nil {
		fmt.Printf("DEBUG: Error occurred while fetching lokasi records: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch lokasi records"})
		return
	}

	// Check if no records were found
	if len(lokasi) == 0 {
		fmt.Println("DEBUG: No records found in table `lokasi`. Returning empty list.")
		c.JSON(http.StatusOK, gin.H{"data": []models.Lokasi{}})
		return
	}

	// Get the total count of records to calculate total pages
	var totalRecords int64
	db.Model(&models.Lokasi{}).Count(&totalRecords)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limitInt)))

	fmt.Printf("DEBUG: Retrieved Lokasi Records: %+v\n", lokasi)
	c.JSON(http.StatusOK, gin.H{
		"data":        lokasi,
		"totalPages":  totalPages,
		"currentPage": pageInt,
	})
}

// Fungsi update untuk Lokasi
func UpdateLokasi(c *gin.Context) {
	// Ambil id dari URL parameter
	id := c.Param("id")
	fmt.Println("Update id:", id) // Debugging ID yang diterima dari URL

	var lokasi models.Lokasi
	db := database.DB

	// Cari lokasi berdasarkan idLokasi
	if err := db.First(&lokasi, "idLokasi = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lokasi not found"})
		return
	}

	// Log data lokasi yang ditemukan
	fmt.Println("Lokasi data found:", lokasi) // Debugging data lokasi yang ditemukan

	// Bind data JSON ke model hanya untuk namaLokasi dan mapLokasi
	var updateData struct {
		NamaLokasi string `json:"namaLokasi"`
		MapLokasi  string `json:"mapLokasi"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log data yang diterima dari request JSON
	fmt.Println("Received update data:", updateData) // Debugging data yang diterima dari frontend

	// Update hanya kolom namaLokasi dan mapLokasi
	if updateData.NamaLokasi != "" {
		lokasi.NamaLokasi = updateData.NamaLokasi
	}
	if updateData.MapLokasi != "" {
		lokasi.MapLokasi = updateData.MapLokasi
	}

	// Simpan perubahan
	if err := db.Save(&lokasi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lokasi"})
		return
	}

	// Kembalikan response sukses
	c.JSON(http.StatusOK, gin.H{"message": "Lokasi updated successfully", "data": lokasi})
}

// DeleteLokasi deletes a specific Lokasi by ID
func DeleteLokasi(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("delete id:", id)
	db := database.DB
	if err := db.Delete(&models.Lokasi{}, "idLokasi = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete lokasi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lokasi deleted successfully"})
}

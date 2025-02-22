package controllers

import (
	"VisualisasiData/database"
	"VisualisasiData/models"
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateKecelakaan creates a new kecelakaan record
func CreateKecelakaan(c *gin.Context) {
	var kecelakaan models.Kecelakaan
	db := database.DB

	// Debug: Log raw JSON body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		log.Println("Error reading request body:", err)
		return
	}
	log.Println("Raw JSON received:", string(body))

	// Bind JSON to struct
	if err := json.Unmarshal(body, &kecelakaan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		log.Println("Error unmarshaling JSON:", err)
		return
	}
	log.Printf("Parsed data: %+v\n", kecelakaan)

	// Validate and parse the waktuInput to proper time format
	if kecelakaan.Waktu == nil || kecelakaan.Waktu.IsZero() {
		// If the waktu is zero or nil, set it to a default value (current time)
		kecelakaan.Waktu = new(time.Time)
		*kecelakaan.Waktu = time.Now() // Assign the address of current time
	}

	// Insert new kecelakaan record
	if err := db.Create(&kecelakaan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create kecelakaan"})
		log.Println("Database insertion error:", err)
		return
	}

	// Respond with the created kecelakaan
	c.JSON(http.StatusOK, gin.H{"message": "Kecelakaan created successfully", "data": kecelakaan})
	log.Println("Kecelakaan created successfully:", kecelakaan)
}

// GetKecelakaan retrieves a kecelakaan record by ID
func GetKecelakaan(c *gin.Context) {
	id := c.Param("id")
	var kecelakaan models.Kecelakaan
	db := database.DB
	// Find kecelakaan by ID
	if err := db.First(&kecelakaan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kecelakaan not found"})
		return
	}

	// Respond with the found kecelakaan
	c.JSON(http.StatusOK, gin.H{"data": kecelakaan})
}

// GetAllKecelakaan retrieves all kecelakaan records
func GetAllKecelakaan(c *gin.Context) {
	var kecelakaan []models.Kecelakaan
	db := database.DB

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
	log.Printf("Pagination offset: %d, limit: %d", offset, limitInt)

	// Adjusting the query to fetch the required fields
	query := db.Model(&models.Kecelakaan{}).
		Select("idKecelakaan", "idKendaraan", "penyebab", "korban", "tanggal", "waktu", "lokasiKecelakaan", "idLokasi", "jenisJalur").
		Where("DeletedAt IS NULL").
		Limit(limitInt).
		Offset(offset)

	// Optionally filter by 'lokasiKecelakaan' query parameter
	lokasiKecelakaan := c.Query("lokasiKecelakaan")
	if lokasiKecelakaan != "" {
		query = query.Where("lokasiKecelakaan LIKE ?", "%"+lokasiKecelakaan+"%")
		log.Printf("Filtering by lokasiKecelakaan: %s", lokasiKecelakaan)
	}

	// Execute the query
	if err := query.Find(&kecelakaan).Error; err != nil {
		log.Printf("Error fetching all kecelakaan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	// Debug log to inspect the retrieved data
	log.Printf("Retrieved kecelakaan data: %v", kecelakaan)

	// Check if no records were found
	if len(kecelakaan) == 0 {
		log.Println("No kecelakaan records found.")
		c.JSON(http.StatusOK, gin.H{"data": []models.Kecelakaan{}})
		return
	}

	// Get the total count of kecelakaan records to calculate total pages
	var totalRecords int64
	if err := db.Model(&models.Kecelakaan{}).Where("DeletedAt IS NULL").Count(&totalRecords).Error; err != nil {
		log.Printf("Error counting total records: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting records"})
		return
	}

	log.Printf("Total records: %d", totalRecords)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limitInt)))

	// Return the data along with pagination info
	c.JSON(http.StatusOK, gin.H{
		"data":        kecelakaan,
		"totalPages":  totalPages,
		"currentPage": pageInt,
	})
}

// UpdateKecelakaan updates a kecelakaan record by ID
func UpdateKecelakaan(c *gin.Context) {
	id := c.Param("id")
	var kecelakaan models.Kecelakaan
	db := database.DB
	// Find kecelakaan by ID
	if err := db.First(&kecelakaan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kecelakaan not found"})
		return
	}

	// Bind the new data from the request body
	if err := c.ShouldBindJSON(&kecelakaan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update kecelakaan record
	if err := db.Save(&kecelakaan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update kecelakaan"})
		return
	}

	// Respond with the updated kecelakaan
	c.JSON(http.StatusOK, gin.H{"message": "Kecelakaan updated successfully", "data": kecelakaan})
}

// DeleteKecelakaan deletes a kecelakaan record by ID
func DeleteKecelakaan(c *gin.Context) {
	id := c.Param("id") // Get ID from URL parameters
	var kecelakaan models.Kecelakaan
	db := database.DB

	// Check if the kecelakaan exists
	if err := db.First(&kecelakaan, id).Error; err != nil {
		// Record not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Kecelakaan not found"})
		log.Printf("Kecelakaan with ID %s not found: %v\n", id, err)
		return
	}

	// Attempt to delete the record
	if err := db.Delete(&kecelakaan).Error; err != nil {
		// Deletion error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete kecelakaan"})
		log.Printf("Failed to delete Kecelakaan with ID %s: %v\n", id, err)
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Kecelakaan deleted successfully"})
	log.Printf("Kecelakaan with ID %s deleted successfully\n", id)
}

func GetDatakecelakaanTahunini(c *gin.Context) {
	var kecelakaan []models.Kecelakaan
	db := database.DB

	// Get the current year
	currentYear := time.Now().Year()

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
	log.Printf("Pagination offset: %d, limit: %d", offset, limitInt)

	// Adjust the query to filter by the current year and optionally by 'lokasiKecelakaan'
	query := db.Model(&models.Kecelakaan{}).
		Select("kecelakaan.lokasiKecelakaan", "kecelakaan.tanggal").
		Joins("INNER JOIN kendaraan ON kendaraan.idKendaraan = kecelakaan.idKendaraan").
		Joins("INNER JOIN lokasi ON lokasi.idLokasi = kecelakaan.idLokasi").
		Where("kecelakaan.DeletedAt IS NULL AND YEAR(kecelakaan.tanggal) = ?", currentYear).
		Limit(limitInt).
		Offset(offset)

	// Optionally filter by 'lokasiKecelakaan' query parameter
	lokasiKecelakaan := c.Query("lokasiKecelakaan")
	if lokasiKecelakaan != "" {
		query = query.Where("kecelakaan.lokasiKecelakaan LIKE ?", "%"+lokasiKecelakaan+"%")
		log.Printf("Filtering by lokasiKecelakaan: %s", lokasiKecelakaan)
	}

	// Execute the query
	if err := query.Find(&kecelakaan).Error; err != nil {
		log.Printf("Error fetching kecelakaan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	// Check if no records were found
	if len(kecelakaan) == 0 {
		log.Println("No kecelakaan records found.")
		c.JSON(http.StatusOK, gin.H{"data": []models.Kecelakaan{}})
		return
	}

	// Get the total count of kecelakaan records for the current year
	var totalRecords int64
	if err := db.Model(&models.Kecelakaan{}).
		Where("DeletedAt IS NULL AND YEAR(tanggal) = ?", currentYear).
		Count(&totalRecords).Error; err != nil {
		log.Printf("Error counting total records: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting records"})
		return
	}

	log.Printf("Total records: %d", totalRecords)

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limitInt)))

	// Return the data along with pagination info
	c.JSON(http.StatusOK, gin.H{
		"data":        kecelakaan,
		"totalPages":  totalPages,
		"currentPage": pageInt,
	})
}

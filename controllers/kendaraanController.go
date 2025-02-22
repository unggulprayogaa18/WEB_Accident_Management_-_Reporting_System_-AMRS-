package controllers

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"VisualisasiData/database"
	"VisualisasiData/models"

	"github.com/gin-gonic/gin"
)

// GetAllKendaraan retrieves all kendaraan records where deleted_at is NULL
// GetAllKendaraan retrieves all kendaraan records where deleted_at is NULL with pagination
func GetAllKendaraan(c *gin.Context) {
	var kendaraan []models.Kendaraan
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

	query := db.Model(&models.Kendaraan{}).
		Select("idKendaraan", "namaKendaraan", "warna", "tipe", "platNomor").
		Where("DeletedAt IS NULL").
		Limit(limitInt).
		Offset(offset)

	// Optionally filter by 'namaKendaraan' query parameter
	namaKendaraan := c.Query("namaKendaraan")
	if namaKendaraan != "" {
		query = query.Where("namaKendaraan LIKE ?", "%"+namaKendaraan+"%")
		log.Printf("Filtering by namaKendaraan: %s", namaKendaraan)
	}

	// Execute the query
	if err := query.Find(&kendaraan).Error; err != nil {
		log.Printf("Error fetching all kendaraan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	// Debug log to inspect the retrieved data
	log.Printf("Retrieved kendaraan data: %v", kendaraan)

	// Check if no records were found
	if len(kendaraan) == 0 {
		log.Println("No kendaraan records found.")
		c.JSON(http.StatusOK, gin.H{"data": []models.Kendaraan{}})
		return
	}

	// Get the total count of kendaraan records to calculate total pages
	var totalRecords int64
	if err := db.Model(&models.Kendaraan{}).Where("DeletedAt IS NULL").Count(&totalRecords).Error; err != nil {
		log.Printf("Error counting total records: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting records"})
		return
	}

	log.Printf("Total records: %d", totalRecords)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limitInt)))

	// Return the data along with pagination info
	c.JSON(http.StatusOK, gin.H{
		"data":        kendaraan,
		"totalPages":  totalPages,
		"currentPage": pageInt,
	})
}

// GetKendaraan retrieves a specific kendaraan record by ID
func GetKendaraan(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Fetching kendaraan with ID: %s", id)
	db := database.DB
	var kendaraan models.Kendaraan
	if err := db.Where("idKendaraan = ? AND DeletedAt IS NULL", id).First(&kendaraan).Error; err != nil {
		log.Printf("Error retrieving kendaraan with ID %s: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Kendaraan not found"})
		return
	}
	log.Printf("Successfully retrieved kendaraan with ID: %s", id)
	c.JSON(http.StatusOK, kendaraan)
}

// CreateKendaraan creates a new kendaraan record
func CreateKendaraan(c *gin.Context) {
	var kendaraan models.Kendaraan
	db := database.DB

	// Bind the JSON request body to the kendaraan struct
	if err := c.ShouldBindJSON(&kendaraan); err != nil {
		log.Printf("Error binding JSON data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Debug log to check if data is received correctly
	log.Printf("Received request to create kendaraan: %+v", kendaraan)

	// Insert the data into the database using a parameterized query
	query := "INSERT INTO kendaraan ( namaKendaraan, warna, tipe, platNomor) VALUES ( ?, ?, ?, ?)"
	result := db.Exec(query, kendaraan.NamaKendaraan, kendaraan.Warna, kendaraan.Tipe, kendaraan.PlatNomor)
	if result.Error != nil {
		log.Printf("Error inserting kendaraan: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating kendaraan"})
		return
	}

	// Send a success response with the created kendaraan data
	c.JSON(http.StatusCreated, kendaraan)
}
func UpdateKendaraan(c *gin.Context) {
	var updatedKendaraan models.Kendaraan
	db := database.DB

	// Get the ID from the URL parameter
	idKendaraan := c.Param("id") // This will get the :id from the URL

	// Log the received ID
	log.Printf("Received ID from URL: %s", idKendaraan)

	// Bind the JSON input from the client to the updatedKendaraan struct
	if err := c.ShouldBindJSON(&updatedKendaraan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Log the parsed data
	log.Printf("Updated data: %+v", updatedKendaraan)

	// Find the existing kendaraan by the ID from the URL
	var kendaraan models.Kendaraan
	if err := db.Where("idKendaraan = ?", idKendaraan).First(&kendaraan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kendaraan not found"})
		return
	}

	// Update the vehicle details
	kendaraan.NamaKendaraan = updatedKendaraan.NamaKendaraan
	kendaraan.Warna = updatedKendaraan.Warna
	kendaraan.Tipe = updatedKendaraan.Tipe
	kendaraan.PlatNomor = updatedKendaraan.PlatNomor

	if err := db.Save(&kendaraan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update kendaraan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// DeleteKendaraan soft-deletes a kendaraan record by ID (sets deleted_at timestamp)
func DeleteKendaraan(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Fetching kendaraan with ID for soft delete: %s", id)

	var kendaraan models.Kendaraan
	db := database.DB
	if err := db.Where("idKendaraan = ? AND DeletedAt IS NULL", id).First(&kendaraan).Error; err != nil {
		log.Printf("Error finding kendaraan with ID %s: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Kendaraan not found"})
		return
	}

	// Mark the record as deleted by setting deleted_at to the current time
	now := time.Now()
	kendaraan.DeletedAt = &now

	log.Printf("Soft-deleting kendaraan with ID %s", id)
	if err := db.Save(&kendaraan).Error; err != nil {
		log.Printf("Error soft-deleting kendaraan with ID %s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting kendaraan"})
		return
	}
	log.Printf("Successfully soft-deleted kendaraan with ID %s", id)
	c.JSON(http.StatusOK, gin.H{"message": "Kendaraan soft-deleted successfully"})
}

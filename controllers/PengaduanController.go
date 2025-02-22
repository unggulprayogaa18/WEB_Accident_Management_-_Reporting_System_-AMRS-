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

// CreatePengaduan creates a new pengaduan record using direct SQL query
func CreatePengaduan(c *gin.Context) {
	// Updated struct with correct field names and JSON tags
	var pengaduan struct {
		TanggalWaktu       time.Time `json:"tanggal_waktu"`
		LokasiKecelakaan   string    `json:"lokasi_kecelakaan"`
		IDKendaraan        string    `json:"id_kendaraan"`
		JumlahKendaraan    uint      `json:"jumlah_kendaraan"`
		Cuaca              string    `json:"cuaca"`
		JalurTertutupTotal string    `json:"jalur_tertutup_total"`
		IDLokasi           string    `json:"lokasiPeruas"`
		JenisJalur         string    `json:"jenisJalur"`
		StatusPengaduan    string    `json:"status_pengaduan"`
	}

	db := database.DB

	// Read raw JSON body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		log.Println("Error reading request body:", err)
		return
	}

	// Log the raw JSON to inspect the incoming request
	log.Println("Raw JSON received:", string(body))

	// Unmarshal JSON data into pengaduan struct
	if err := json.Unmarshal(body, &pengaduan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		log.Println("Error unmarshaling JSON:", err)
		return
	}

	// Validate required fields
	if pengaduan.TanggalWaktu.IsZero() || pengaduan.LokasiKecelakaan == "" || pengaduan.JumlahKendaraan == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		log.Println("Validation error: Missing required fields")
		return
	}

	// Prepare the SQL query using the correct column names.
	query := `
        INSERT INTO Pengaduan (
            tanggal_waktu, 
            lokasi_kecelakaan, 
            id_kendaraan, 
            jumlah_kendaraan, 
            id_lokasi, 
            jenis_jalur, 
            cuaca, 
            jalur_tertutup_total, 
            status_pengaduan
        )
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	// Execute the SQL query with the data.
	// IMPORTANT: Make sure the parameters are in the same order as in your query.
	if err := db.Exec(query,
		pengaduan.TanggalWaktu,
		pengaduan.LokasiKecelakaan,
		pengaduan.IDKendaraan,
		pengaduan.JumlahKendaraan,
		pengaduan.IDLokasi,
		pengaduan.JenisJalur,
		pengaduan.Cuaca,
		pengaduan.JalurTertutupTotal, // Using the correct field name
		pengaduan.StatusPengaduan,
	).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pengaduan"})
		log.Println("Database insertion error:", err)
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Pengaduan created successfully"})
	log.Println("Pengaduan created successfully")
}

// GetAllPengaduan retrieves all pengaduan records
func GetAllPengaduan(c *gin.Context) {
	var pengaduan []models.Pengaduan
	db := database.DB

	// Ambil parameter page dan limit dari query string
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "7")
	log.Printf("Received query parameters - page: %s, limit: %s", page, limit)

	// Konversi parameter page dan limit ke integer
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Printf("Error converting page parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Printf("Error converting limit parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	// Hitung offset untuk pagination
	offset := (pageInt - 1) * limitInt
	log.Printf("Calculated offset: %d (page: %d, limit: %d)", offset, pageInt, limitInt)

	// Bangun query dengan memilih kolom-kolom yang diperlukan
	query := db.Model(&models.Pengaduan{}).
		Select("id_pengaduan", "tanggal_waktu", "lokasi_kecelakaan", "id_kendaraan", "jumlah_kendaraan", "id_lokasi", "jenis_jalur", "cuaca", "jalur_tertutup_total", "status_pengaduan").
		Limit(limitInt).
		Offset(offset)

	// Filter opsional berdasarkan 'lokasi_kecelakaan'
	lokasiKecelakaan := c.Query("lokasi_kecelakaan")
	if lokasiKecelakaan != "" {
		query = query.Where("lokasi_kecelakaan LIKE ?", "%"+lokasiKecelakaan+"%")
		log.Printf("Applying filter on lokasi_kecelakaan: %s", lokasiKecelakaan)
	}

	// Eksekusi query
	if err := query.Find(&pengaduan).Error; err != nil {
		log.Printf("Error fetching pengaduan data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	// Log data yang telah diambil
	log.Printf("Fetched %d pengaduan records", len(pengaduan))
	if len(pengaduan) == 0 {
		log.Println("No pengaduan records found for the given query parameters.")
		c.JSON(http.StatusOK, gin.H{"data": []models.Pengaduan{}})
		return
	}

	// Ambil total jumlah record untuk perhitungan total pages (tanpa pagination)
	var totalRecords int64
	if err := db.Model(&models.Pengaduan{}).Count(&totalRecords).Error; err != nil {
		log.Printf("Error counting total records: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting records"})
		return
	}
	log.Printf("Total records in database: %d", totalRecords)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limitInt)))
	log.Printf("Calculated total pages: %d", totalPages)

	// Log sebelum mengirim respons akhir
	log.Printf("Sending response: currentPage=%d, totalPages=%d, records on current page=%d", pageInt, totalPages, len(pengaduan))
	c.JSON(http.StatusOK, gin.H{
		"currentPage": pageInt,
		"data":        pengaduan,
		"totalPages":  totalPages,
		"message":     "Pengaduan berhasil diperbarui",
	})

}

func UpdateStatusValid(c *gin.Context) {
	idPengaduan := c.Param("idPengaduan")

	// Retrieve the pengaduan record by id
	var pengaduan models.Pengaduan
	if err := database.DB.First(&pengaduan, "id_pengaduan = ?", idPengaduan).Error; err != nil {
		log.Println("Error fetching pengaduan:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengaduan not found"})
		return
	}

	// Update the status of pengaduan
	pengaduan.StatusPengaduan = "valid" // Correct field name here
	if err := database.DB.Save(&pengaduan).Error; err != nil {
		log.Println("Error updating pengaduan status:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pengaduan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengaduan berhasil diperbarui"})
}

// UpdateStatusTidakValid untuk mengubah status pengaduan menjadi tidak valid
func UpdateStatusTidakValid(c *gin.Context) {
	idPengaduan := c.Param("idPengaduan")

	// Retrieve the pengaduan record by id
	var pengaduan models.Pengaduan
	if err := database.DB.First(&pengaduan, "id_pengaduan = ?", idPengaduan).Error; err != nil {
		log.Println("Error fetching pengaduan:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengaduan not found"})
		return
	}

	// Update the status of pengaduan
	pengaduan.StatusPengaduan = "tidak_valid" // Correct field name here
	if err := database.DB.Save(&pengaduan).Error; err != nil {
		log.Println("Error updating pengaduan status:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pengaduan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengaduan berhasil diperbarui"})
}

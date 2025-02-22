package controllers

import (
	"net/http"
	"strconv"

	"VisualisasiData/database"
	"VisualisasiData/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPengaduan(c *gin.Context) {
	// Mengambil parameter query "limit" dan "offset"
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	// Konversi string ke integer dengan default value jika error
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // default limit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0 // default offset
	}

	var pengaduan []models.Pengaduan
	db := database.DB

	// Membangun query dengan field yang dipilih, limit, dan offset
	if err := db.Model(&models.Pengaduan{}).
		Select("id_pengaduan", "tanggal_waktu", "lokasi_kecelakaan", "id_kendaraan", "jumlah_kendaraan", "id_lokasi", "jenis_jalur", "cuaca", "jalur_tertutup_total", "status_pengaduan").
		Limit(limit).
		Offset(offset).
		Find(&pengaduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	totalPages := 1 // Anda perlu menghitung total pages berdasarkan data

	c.JSON(http.StatusOK, gin.H{
		"data":       pengaduan,
		"totalPages": totalPages,
	})
}

func DeletePenangan(c *gin.Context, db *gorm.DB) {
	// Ambil ID dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Konversi ke integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Cek apakah data dengan ID tersebut ada di tabel pengaduan_kecelakaan
	var pengaduan models.Pengaduan
	if err := db.First(&pengaduan, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pengaduan not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Pengaduan"})
		return
	}

	// Lakukan soft delete
	if err := db.Delete(&pengaduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Pengaduan"})
		return
	}

	// Berikan respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Pengaduan soft-deleted successfully"})
}

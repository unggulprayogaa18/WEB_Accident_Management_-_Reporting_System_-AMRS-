package controllers

import (
	"VisualisasiData/database"
	"VisualisasiData/models"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetFormData(c *gin.Context) {
	var kendaraan []models.Kendaraan
	var lokasi []models.Lokasi
	db := database.DB

	// Check if the database connection is valid
	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection is not initialized"})
		return
	}

	// Fetch data for Kendaraan
	if err := db.Find(&kendaraan).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch kendaraan data"})
		return
	}

	// Fetch data for Lokasi
	if err := db.Find(&lokasi).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch lokasi data"})
		return
	}

	// Log the fetched data for debugging purposes
	log.Println("Kendaraan:", kendaraan)
	log.Println("Lokasi:", lokasi)

	// Menyusun data kendaraan dengan ID dan NamaKendaraan
	kendaraanData := make([]map[string]string, len(kendaraan))
	for i, k := range kendaraan {
		kendaraanData[i] = map[string]string{
			"idKendaraan":   fmt.Sprintf("%d", k.IdKendaraan),
			"namaKendaraan": k.NamaKendaraan,
		}
	}

	// Menyusun data lokasi dengan ID dan NamaLokasi
	lokasiData := make([]map[string]string, len(lokasi))
	for i, l := range lokasi {
		lokasiData[i] = map[string]string{
			"idLokasi":   l.IdLokasi,
			"namaLokasi": l.NamaLokasi,
		}
	}

	// Kembalikan data dalam format JSON
	c.JSON(200, gin.H{
		"kendaraan": kendaraanData,
		"lokasi":    lokasiData,
	})
}

func HitungTotalKecelakaan(c *gin.Context) {
	var total int64
	db := database.DB

	// Filter kecelakaan berdasarkan tahun saat ini
	currentYear := time.Now().Year()
	if err := db.Model(&models.Kecelakaan{}).
		Where("YEAR(tanggal) = ?", currentYear).
		Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching total kecelakaan"})
		return
	}
	log.Printf("Current year: %d, Total kecelakaan: %d", currentYear, total)

	c.JSON(http.StatusOK, gin.H{"total": total})
}

func Getdatakecelakaansemuatahun(c *gin.Context) {
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

	// Adjust the query to group by 'lokasiKecelakaan' and month of 'tanggal', and sum 'jumlahKecelakaan'
	query := db.Model(&models.Kecelakaan{}).
		Select("kecelakaan.lokasiKecelakaan", "kecelakaan.tanggal").
		Joins("INNER JOIN kendaraan ON kendaraan.idKendaraan = kecelakaan.idKendaraan").
		Joins("INNER JOIN lokasi ON lokasi.idLokasi = kecelakaan.idLokasi").
		Where("kecelakaan.DeletedAt IS NULL").
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
		"kecelakaan_data": kecelakaan, // Ubah key untuk menghindari konflik
		"total_pages":     totalPages,
		"current_page":    pageInt,
	})

}

func HitungTotalKecelakaanMobileTahunIni(c *gin.Context) {
	var result []map[string]interface{}
	db := database.DB

	// Filter kecelakaan berdasarkan tahun saat ini dan menghitung kecelakaan berdasarkan tipe kendaraan
	currentYear := time.Now().Year()

	// Query untuk menghitung jumlah kecelakaan per tipe kendaraan
	if err := db.Model(&models.Kecelakaan{}).
		Select("kendaraan.tipe AS tipeKendaraan, COUNT(*) AS total").
		Joins("JOIN kendaraan ON kendaraan.idKendaraan = kecelakaan.idKendaraan").
		Where("YEAR(kecelakaan.tanggal) = ?", currentYear).
		Group("kendaraan.tipe").
		Order("total DESC"). // Mengurutkan berdasarkan jumlah kecelakaan terbanyak
		Limit(1).            // Membatasi hanya satu tipe dengan jumlah kecelakaan tertinggi
		Find(&result).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching total kecelakaan berdasarkan tipe kendaraan"})
		return
	}

	// Mengecek apakah ada data yang ditemukan
	if len(result) > 0 {
		// Menampilkan tipe kendaraan dengan kecelakaan terbanyak
		tipeKendaraan := result[0]["tipeKendaraan"]
		total := result[0]["total"]
		log.Printf("Tipe kendaraan dengan kecelakaan terbanyak: %v, Total: %v", tipeKendaraan, total)
		c.JSON(http.StatusOK, gin.H{"tipeKendaraan": tipeKendaraan, "total": total})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Tidak ada data kecelakaan"})
	}
}

func HitungTotalKecelakaanPerLokasi(c *gin.Context) {
	db := database.DB

	// Ambil tahun saat ini
	currentYear := time.Now().Year()

	// Menampung hasil query menggunakan slice of maps
	var hasil []map[string]interface{}

	// Query untuk menghitung total kecelakaan per lokasi berdasarkan tahun
	if err := db.Table("kecelakaan").
		Select("lokasi.namaLokasi AS namaLokasi, COUNT(kecelakaan.idKecelakaan) AS totalKecelakaan").
		Joins("JOIN lokasi ON lokasi.idLokasi = kecelakaan.idLokasi").
		Where("YEAR(kecelakaan.tanggal) = ?", currentYear).
		Group("lokasi.namaLokasi").
		Scan(&hasil).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	// Kembalikan hasil dalam format JSON
	c.JSON(http.StatusOK, hasil)
}

func HitungPenyebabTertinggi(c *gin.Context) {
	var penyebabCount []struct {
		Penyebab string
		Count    int64
	}
	db := database.DB

	// Ambil data penyebab dari tabel Kecelakaan dan hitung frekuensinya
	if err := db.Model(&models.Kecelakaan{}).
		Select("penyebab, COUNT(*) as count").
		Group("penyebab").
		Order("count DESC").
		Limit(1).
		Find(&penyebabCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching penyebab kecelakaan"})
		return
	}

	if len(penyebabCount) > 0 {
		log.Printf("Penyebab terbanyak: %s dengan %d kejadian", penyebabCount[0].Penyebab, penyebabCount[0].Count)
		c.JSON(http.StatusOK, gin.H{
			"penyebab": penyebabCount[0].Penyebab,
			"count":    penyebabCount[0].Count,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "No data found"})
	}
}

func KejadianKorban(c *gin.Context) {
	db := database.DB

	// Get the page and limit parameters from the query
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "40")

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

	// Adjusting the query to fetch the required fields with JOIN
	var kecelakaan []map[string]interface{}

	query := db.Table("kecelakaan").
		Select("DATE_FORMAT(kecelakaan.waktu, '%d/%m/%Y %H:%i:%s') as waktu, "+ // Formatted waktu
			"lokasi.namaLokasi, "+
			"kecelakaan.lokasiKecelakaan, "+
			"kecelakaan.idLokasi, "+
			"kecelakaan.penyebab, "+
			"kecelakaan.korban, "+
			"kecelakaan.jenisJalur").
		Joins("INNER JOIN kendaraan ON kendaraan.idKendaraan = kecelakaan.idKendaraan").
		Joins("INNER JOIN lokasi ON lokasi.idLokasi = kecelakaan.idLokasi").
		Where("YEAR(kecelakaan.tanggal) = ?", time.Now().Year()). // Filter by current year
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
		log.Printf("Error fetching kecelakaan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	// Check if no records were found
	if len(kecelakaan) == 0 {
		log.Println("No kecelakaan records found.")
		c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
		return
	}

	// Process the korban column to calculate selamat, meninggal, luka-luka
	for i := range kecelakaan {
		korbanStr, ok := kecelakaan[i]["korban"].(string)
		if ok {
			// Initialize counts
			var selamat, meninggal, lukaLuka int

			// Split the korban string into parts: "2 selamat - 3 meninggal - 4 luka-luka"
			parts := strings.Split(korbanStr, " - ")

			for _, part := range parts {
				// Parse each part to extract the number and the type
				// Example: "2 selamat", "3 meninggal", "4 luka-luka"
				segments := strings.Fields(part)
				if len(segments) == 2 {
					count, err := strconv.Atoi(segments[0])
					if err != nil {
						log.Printf("Error parsing count in korban string: %v", err)
						continue
					}

					switch segments[1] {
					case "selamat":
						selamat += count
					case "meninggal":
						meninggal += count
					case "luka-luka":
						lukaLuka += count
					}
				}
			}

			// Assign the calculated values back to the record
			kecelakaan[i]["selamat"] = selamat
			kecelakaan[i]["meninggal"] = meninggal
			kecelakaan[i]["luka_luka"] = lukaLuka
		}
	}

	// Get the total count of kecelakaan records to calculate total pages
	var totalRecords int64
	if err := db.Model(&models.Kecelakaan{}).Where("YEAR(tanggal) = ?", time.Now().Year()).Count(&totalRecords).Error; err != nil {
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
func JenisJalurPengaduan(c *gin.Context) {
	db := database.DB

	// Query the database to get the counts of accidents for each jenisJalur where status_pengaduan is valid
	rows, err := db.Raw(`
        SELECT jenis_jalur, COUNT(*) AS count
        FROM pengaduan
        WHERE status_pengaduan = 'valid'
        GROUP BY jenis_jalur
    `).Rows()
	if err != nil {
		log.Printf("Error fetching data by jenis_jalur: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}
	defer rows.Close()

	// Initialize counts for each jenisJalur
	jenisJalurCounts := map[string]int{
		"A":          0,
		"B":          0,
		"None Jalur": 0,
	}

	// Log initial counts (before populating)
	log.Printf("Initial Jenis Jalur Counts: %v", jenisJalurCounts)

	// Check if rows are empty or not
	rowCount := 0

	// Populate the map with counts from the query results
	for rows.Next() {
		var jenisJalur string
		var count int
		if err := rows.Scan(&jenisJalur, &count); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
			return
		}

		log.Printf("Fetched jenis_jalur: %s, Count: %d", jenisJalur, count)

		// Update the count for the corresponding jenisJalur
		if jenisJalur == "A" || jenisJalur == "B" {
			jenisJalurCounts[jenisJalur] = count
		} else {
			jenisJalurCounts["None Jalur"] += count
		}

		rowCount++
	}

	// Log the result after processing rows
	if rowCount == 0 {
		log.Printf("No rows returned for jenis_jalur with status_pengaduan = 'valid'.")
	} else {
		log.Printf("Final Jenis Jalur Counts: %v", jenisJalurCounts)
	}

	// Send the result as JSON
	c.JSON(http.StatusOK, gin.H{"data": jenisJalurCounts})
}

func CountTopLocationpengaduan(c *gin.Context) {
	db := database.DB

	// Query untuk menghitung jumlah kecelakaan berdasarkan lokasi
	rows, err := db.Raw(`
        SELECT lokasi_kecelakaan, COUNT(*) AS count
        FROM pengaduan
        WHERE status_pengaduan = 'valid'
        GROUP BY lokasi_kecelakaan
        ORDER BY count DESC
    `).Rows()
	if err != nil {
		log.Printf("Error fetching top locations: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}
	defer rows.Close()

	// Data map to store locations and their counts
	data := make(map[string]int)

	// Process query results and populate the map
	for rows.Next() {
		var lokasi string
		var count int
		if err := rows.Scan(&lokasi, &count); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
			return
		}
		// Store the count for each lokasi
		data[lokasi] = count
	}

	// Log data for debugging purposes
	log.Printf("Top Locations Data: %v", data)

	// Send the result as JSON
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

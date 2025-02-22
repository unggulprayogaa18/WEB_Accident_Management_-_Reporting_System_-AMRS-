package controllers

import (
	"VisualisasiData/database"
	"VisualisasiData/models"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func INNERJOINdataKecelakaan(c *gin.Context) {
	var results []map[string]interface{}
	db := database.DB

	// Pagination parameters
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "15")
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
	offset := (pageInt - 1) * limitInt

	// Query with INNER JOIN
	if err := db.Table("kecelakaan").
		Select(`
            kecelakaan.idKecelakaan, 
            kecelakaan.penyebab, 
            kecelakaan.korban, 
            kecelakaan.tanggal, 
            kecelakaan.waktu,
            kecelakaan.jenisJalur,
            kecelakaan.lokasiKecelakaan,
            kendaraan.warna, 
            kendaraan.platNomor, 
            kendaraan.namaKendaraan,
            lokasi.namaLokasi`).
		Joins("INNER JOIN kendaraan ON kecelakaan.idKendaraan = kendaraan.idKendaraan").
		Joins("INNER JOIN lokasi ON kecelakaan.idLokasi = lokasi.idLokasi").
		Where("kecelakaan.deletedAt IS NULL AND kendaraan.deletedAt IS NULL").
		Limit(limitInt).
		Offset(offset).
		Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	// Total records for pagination
	var totalRecords int64
	if err := db.Table("kecelakaan").
		Joins("INNER JOIN kendaraan ON kecelakaan.idKendaraan = kendaraan.idKendaraan").
		Joins("INNER JOIN lokasi ON kecelakaan.idLokasi = lokasi.idLokasi").
		Where("kecelakaan.deletedAt IS NULL").
		Count(&totalRecords).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count records"})
		return
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limitInt)))

	// Return results with pagination
	c.JSON(http.StatusOK, gin.H{
		"data":        results,
		"totalPages":  totalPages,
		"currentPage": pageInt,
	})
}

// // CountDataByMonth counts the number of records for each month in the current year
func CountDataByYearAndMonth(c *gin.Context) {
	// Tentukan tahun awal dan tahun saat ini
	startYear := 2023
	currentYear := time.Now().Year()
	db := database.DB

	// Inisialisasi map nested untuk menyimpan data per tahun dan per bulan
	yearMonthCounts := make(map[int]map[string]int)

	// Daftar nama bulan
	months := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	// Inisialisasi setiap tahun dengan semua bulan bernilai 0
	for year := startYear; year <= currentYear; year++ {
		yearMonthCounts[year] = make(map[string]int)
		for _, month := range months {
			yearMonthCounts[year][month] = 0
		}
	}

	// Query dengan UNION ALL tanpa filter berdasarkan bulan saat ini,
	// melainkan mengambil semua data dari tahun startYear sampai currentYear.
	query := `
        SELECT t.year, t.month, SUM(t.count) AS count
        FROM (
            SELECT YEAR(k.tanggal) AS year, MONTH(k.tanggal) AS month, COUNT(*) AS count
            FROM kecelakaan k
            WHERE YEAR(k.tanggal) BETWEEN ? AND ?
            GROUP BY YEAR(k.tanggal), MONTH(k.tanggal)
            UNION ALL
            SELECT YEAR(p.tanggal_waktu) AS year, MONTH(p.tanggal_waktu) AS month, COUNT(*) AS count
            FROM pengaduan p
            WHERE YEAR(p.tanggal_waktu) BETWEEN ? AND ?
              AND p.status_pengaduan = 'valid'
            GROUP BY YEAR(p.tanggal_waktu), MONTH(p.tanggal_waktu)
        ) t
        GROUP BY t.year, t.month
    `

	rows, err := db.Raw(query, startYear, currentYear, startYear, currentYear).Rows()
	if err != nil {
		log.Printf("Error fetching data by year and month: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}
	defer rows.Close()

	// Proses hasil query dan masukkan ke dalam map
	for rows.Next() {
		var year, month, count int
		if err := rows.Scan(&year, &month, &count); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
			return
		}
		log.Printf("Fetched year: %d, month: %d, count: %d", year, month, count)

		// Konversi nomor bulan ke nama bulan (misal: 2 -> February)
		monthName := months[month-1]
		yearMonthCounts[year][monthName] = count
	}

	// Kirim hasil sebagai JSON
	c.JSON(http.StatusOK, gin.H{
		"startYear": startYear,
		"endYear":   currentYear,
		"data":      yearMonthCounts,
	})
}

// func CountDataByMonth(c *gin.Context) {
// 	db := database.DB

// 	// Get the current year
// 	currentYear := time.Now().Year()

// 	// Initialize a map to store the counts for each month
// 	monthCounts := make(map[string]int)

// 	// Define month names
// 	months := []string{
// 		"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",
// 	}

// 	// Initialize the map with 0 counts for each month
// 	for _, month := range months {
// 		monthCounts[month] = 0
// 	}

// 	// Query the database to get all records for the current year
// 	rows, err := db.Raw(`
//         SELECT MONTH(tanggal) AS month, COUNT(*) AS count
//         FROM kecelakaan
//         WHERE YEAR(tanggal) = ? AND DeletedAt IS NULL
//         GROUP BY MONTH(tanggal)
//     `, currentYear).Rows()
// 	if err != nil {
// 		log.Printf("Error fetching data by month: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
// 		return
// 	}
// 	defer rows.Close()

// 	// Populate the map with counts from the query results
// 	for rows.Next() {
// 		var month int
// 		var count int
// 		if err := rows.Scan(&month, &count); err != nil {
// 			log.Printf("Error scanning row: %v", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
// 			return
// 		}

// 		// Convert numeric month to month name
// 		monthName := months[month-1]
// 		monthCounts[monthName] = count
// 	}

//		// Send the result as JSON
//		c.JSON(http.StatusOK, gin.H{"year": currentYear, "data": monthCounts})
//	}
//
// waktuKejadianKecelakaan retrieves all kecelakaan records along with kendaraan details and count per time
func KejadianWaktu(c *gin.Context) {
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
		Select("DATE_FORMAT(kecelakaan.waktu, '%H:%i:%s') as waktu, "+
			"COUNT(kecelakaan.idKecelakaan) as totalKecelakaan, "+
			"kendaraan.namaKendaraan, kendaraan.tipe as tipeKendaraan, "+
			"kendaraan.warna as warnaKendaraan, kendaraan.platNomor as platNomorKendaraan, "+
			"lokasi.namaLokasi,  kecelakaan.penyebab, kecelakaan.korban, "+
			"kecelakaan.lokasiKecelakaan , kecelakaan.jenisJalur").
		Joins("INNER JOIN kendaraan ON kendaraan.idKendaraan = kecelakaan.idKendaraan").
		Joins("INNER JOIN lokasi ON lokasi.idLokasi = kecelakaan.idLokasi").
		Where("YEAR(kecelakaan.tanggal) = ?", time.Now().Year()). // Filter by current year
		Group("waktu, kendaraan.idKendaraan, lokasi.idLokasi, kecelakaan.penyebab, kecelakaan.korban, kecelakaan.lokasiKecelakaan").
		Limit(limitInt)

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

func CountDataByJenisJalur(c *gin.Context) {
	db := database.DB

	// Query the database to get the counts of accidents for each jenisJalur without date filtering
	rows, err := db.Raw(`
        SELECT jenisJalur, COUNT(*) AS count
        FROM kecelakaan
        WHERE DeletedAt IS NULL
        GROUP BY jenisJalur
    `).Rows()
	if err != nil {
		log.Printf("Error fetching data by jenisJalur: %v", err)
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

	// Populate the map with counts from the query results
	for rows.Next() {
		var jenisJalur string
		var count int
		if err := rows.Scan(&jenisJalur, &count); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
			return
		}

		// Update the count for the corresponding jenisJalur
		jenisJalurCounts[jenisJalur] = count
	}

	// Send the result as JSON
	c.JSON(http.StatusOK, gin.H{"data": jenisJalurCounts})
}

func CountTopLocation(c *gin.Context) {
	db := database.DB

	// Query untuk menghitung jumlah kecelakaan berdasarkan lokasi
	rows, err := db.Raw(`
        SELECT lokasiKecelakaan, COUNT(*) AS count
        FROM kecelakaan
        WHERE DeletedAt IS NULL
        GROUP BY lokasiKecelakaan
        ORDER BY count DESC
    `).Rows()
	if err != nil {
		log.Printf("Error fetching top locations: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}
	defer rows.Close()

	data := make(map[string]int)

	// Ambil semua data dari hasil query
	for rows.Next() {
		var lokasi string
		var count int
		if err := rows.Scan(&lokasi, &count); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
			return
		}
		data[lokasi] = count
	}

	// Kirim hasil sebagai JSON
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

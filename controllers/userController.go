package controllers

import (
	"VisualisasiData/database"
	"VisualisasiData/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Fungsi untuk login user dengan password dibandingkan secara langsung (plain text)

// Fungsi untuk membuat user baru dengan password yang disimpan dalam bentuk plain text
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Simpan password secara langsung tanpa hashing (plain text)
	// user.Password sudah dalam bentuk plain text, tidak diubah

	// Simpan user ke database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Beri respon sukses setelah user berhasil dibuat
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

// Fungsi untuk login user
func LoginUser(c *gin.Context) {

	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind input JSON
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Cari user berdasarkan email
	var user models.User
	if err := database.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Bandingkan password yang diberikan dengan password yang ada di database (plain text comparison)
	if user.Password != loginData.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Setelah berhasil login, simpan user dalam konteks (untuk keperluan session atau token)
	c.Set("user", user) // Menyimpan data user di context

	// Redirect berdasarkan role (admin atau user)
	if user.Role == "admin" {
		// Jika admin, arahkan ke halaman admin
		c.JSON(http.StatusOK, gin.H{
			"message":  "Login successful",
			"redirect": "/frontend/admin.html", // Halaman frontend untuk admin
		})
	} else if user.Role == "user" {
		// Jika user, arahkan ke halaman user
		c.JSON(http.StatusOK, gin.H{
			"message":  "Login successful",
			"redirect": "/frontend/users.html", // Halaman frontend untuk user
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not recognized"})
	}
}

// fungsi menggunakan bycrpt
// package controllers

// import (
// 	"net/http"
// 	"projectsanjaya/database"
// 	"projectsanjaya/models"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/crypto/bcrypt"
// )

// // Fungsi untuk membuat user baru dengan password yang di-hash
// func CreateUser(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	// Hash password sebelum disimpan
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
// 		return
// 	}
// 	user.Password = string(hashedPassword)

// 	// Simpan user ke database
// 	if err := database.DB.Create(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
// }

// // Fungsi untuk login user
// func LoginUser(c *gin.Context) {
// 	var loginData struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	// Bind input JSON
// 	if err := c.ShouldBindJSON(&loginData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	// Cari user berdasarkan email
// 	var user models.User
// 	if err := database.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
// 		return
// 	}

// 	// Bandingkan password yang diberikan dengan password yang ada di database
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
// 		return
// 	}

// 	// Setelah berhasil login, simpan user dalam konteks (untuk keperluan session atau token)
// 	c.Set("user", user) // Menyimpan data user di context

// 	// Redirect berdasarkan role (admin atau user)
// 	if user.Role == "admin" {
// 		// Jika admin, arahkan ke halaman admin
// 		c.JSON(http.StatusOK, gin.H{
// 			"message":  "Login successful",
// 			"redirect": "/frontend/admin.html", // Halaman frontend untuk admin
// 		})
// 	} else if user.Role == "user" {
// 		// Jika user, arahkan ke halaman user
// 		c.JSON(http.StatusOK, gin.H{
// 			"message":  "Login successful",
// 			"redirect": "/frontend/frontend.html", // Halaman frontend untuk user
// 		})

// 	} else {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not recognized"})
// 	}
// }

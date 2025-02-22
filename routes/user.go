package routes

import (
	"VisualisasiData/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/api/users", controllers.CreateUser) // Endpoint untuk membuat user baru
	r.POST("/api/login", controllers.LoginUser)  // Endpoint untuk login

	r.GET("/Form.html", func(c *gin.Context) {
		c.File("./frontend/Form.html")
	})

	r.GET("/sign-up.html", func(c *gin.Context) {
		c.File("./frontend/sign-up.html")
	})

	r.GET("/admin.html", func(c *gin.Context) {
		c.File("./frontend/admin.html")
	})
	// Route to fetch form data from the database and display in the form

	// // Admin Routes
	// adminGroup := r.Group("/api/admin")
	// adminGroup.Use(AdminMiddleware()) // Middleware untuk memverifikasi role admin
	// {
	// 	// adminGroup.PUT("/users/:id", controllers.UpdateUser)    // Mengupdate user
	// 	// adminGroup.DELETE("/users/:id", controllers.DeleteUser) // Menghapus user
	// }
}

// // Middleware untuk memverifikasi role admin
// func AdminMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		user, exists := c.Get("user") // Ambil user dari context
// 		if !exists {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
// 			c.Abort()
// 			return
// 		}

// 		// Pastikan user memiliki tipe yang benar
// 		if u, ok := user.(models.User); ok {
// 			if u.Role != "admin" {
// 				c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission"})
// 				c.Abort()
// 				return
// 			}
// 		} else {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }

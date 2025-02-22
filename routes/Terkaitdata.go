package routes

import (
	"VisualisasiData/controllers"

	"github.com/gin-gonic/gin"
)

func TerkaitdataRoute(r *gin.Engine) {
	// Enable CORS
	// This will allow all origins, you can customize as needed
	r.GET("/ambildatakendaranlokasi/getformdata", controllers.GetFormData)
	r.GET("/api/kecelakaan/total", controllers.HitungTotalKecelakaan)
	r.GET("/api/chartbulan", controllers.CountDataByYearAndMonth)
	r.GET("/api/waktukejadian", controllers.KejadianWaktu)
	r.GET("/api/KejadianKorban", controllers.KejadianKorban)
	r.GET("/api/countByJenisJalurbypengaduan", controllers.JenisJalurPengaduan)
	r.GET("/api/countByJenisJalur", controllers.CountDataByJenisJalur)
	r.GET("/top-location", controllers.CountTopLocation)
	r.GET("/top-location2", controllers.CountTopLocationpengaduan)

	r.GET("/tables1.html", func(c *gin.Context) {
		c.File("./frontend/tables1.html")
	})
	r.GET("/tables2.html", func(c *gin.Context) {
		c.File("./frontend/tables2.html")
	})
	r.GET("/pengaduan.html", func(c *gin.Context) {
		c.File("./frontend/FormPengaduan.html")
	})
	r.GET("/penanganan.html", func(c *gin.Context) {
		c.File("./frontend/FormPenanganan.html")
	})
	r.GET("/grafik.html", func(c *gin.Context) {
		c.File("./frontend/grafik.html")
	})
	r.GET("/tables3.html", func(c *gin.Context) {
		c.File("./frontend/tables3.html")
	})

	r.GET("/users.html", func(c *gin.Context) {
		c.File("./frontend/users.html")
	})
}

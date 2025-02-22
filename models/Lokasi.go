package models

type Lokasi struct {
	// Gunakan kolom idLokasi sebagai primary key
	IdLokasi   string `json:"idLokasi" gorm:"primaryKey;column:idLokasi;unique;not null"`
	NamaLokasi string `json:"namaLokasi" gorm:"column:namaLokasi;not null"`
	MapLokasi  string `json:"mapLokasi" gorm:"column:mapLokasi;not null"`
}

// TableName overrides the default table name for Lokasi
func (Lokasi) TableName() string {
	return "lokasi" // Explicitly set the table name to "lokasi"
}

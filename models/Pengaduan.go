package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Pengaduan struct {
	IDPengaduan        uint      `json:"id_pengaduan" gorm:"primaryKey;autoIncrement"`
	TanggalWaktu       time.Time `json:"tanggal_waktu"`
	LokasiKecelakaan   string    `json:"lokasi_kecelakaan"`
	IDKendaraan        string    `json:"id_kendaraan"`
	JumlahKendaraan    uint      `json:"jumlah_kendaraan"`
	IDLokasi           string    `json:"id_lokasi"`
	JenisJalur         string    `json:"jenis_jalur" gorm:"type:ENUM('A', 'B')"`
	Cuaca              string    `json:"cuaca"`
	JalurTertutupTotal string    `json:"jalur_tertutup_total" gorm:"type:ENUM('ya', 'tidak')"`
	StatusPengaduan    string    `json:"status_pengaduan" gorm:"column:status_pengaduan;type:ENUM('valid', 'tidak_valid', 'belum_ditanggapi')"`
}

func (Pengaduan) TableName() string {
	return "Pengaduan"
}

// Custom UnmarshalJSON for Pengaduan struct
func (p *Pengaduan) UnmarshalJSON(data []byte) error {
	type Alias Pengaduan
	aux := &struct {
		TanggalWaktu string `json:"tanggal_waktu"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	// Unmarshal JSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Validate date-time format
	if aux.TanggalWaktu != "" {
		parsedDateTime, err := time.Parse("2006-01-02 15:04", aux.TanggalWaktu)
		if err != nil {
			return fmt.Errorf("invalid tanggal_waktu format: %v", err)
		}
		p.TanggalWaktu = parsedDateTime
	}

	return nil
}

package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Kecelakaan struct {
	IdKecelakaan     uint       `json:"idKecelakaan" gorm:"primaryKey;autoIncrement;column:idKecelakaan;not null"`
	IdKendaraan      uint       `json:"idKendaraan" gorm:"column:idKendaraan;not null"`
	Penyebab         string     `json:"penyebab" gorm:"column:penyebab;not null"`
	Korban           string     `json:"korban" gorm:"column:korban;not null"`
	Tanggal          time.Time  `json:"tanggal" gorm:"column:tanggal;not null"`
	Waktu            *time.Time `json:"waktu" gorm:"column:waktu;not null"`
	LokasiKecelakaan string     `json:"lokasiKecelakaan" gorm:"column:lokasiKecelakaan;not null"`
	LokasiPeruas     string     `json:"lokasiPeruas" gorm:"column:idLokasi;not null"`
	JenisJalur       string     `json:"jenisJalur" gorm:"column:jenisJalur;not null"`
}

func (Kecelakaan) TableName() string {
	return "kecelakaan" // Explicitly set the table name to "kecelakaan"
}

// Custom UnmarshalJSON untuk struct Kecelakaan
func (k *Kecelakaan) UnmarshalJSON(data []byte) error {
	type Alias Kecelakaan
	aux := &struct {
		Tanggal    string `json:"tanggal"`
		WaktuInput string `json:"waktuInput"`
		*Alias
	}{
		Alias: (*Alias)(k),
	}

	// Unmarshal data yang masuk ke dalam struct
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Validasi dan konversi string Tanggal ke time.Time
	if aux.Tanggal != "" {
		var err error
		k.Tanggal, err = time.Parse("2006-01-02", aux.Tanggal) // Format "YYYY-MM-DD"
		if err != nil {
			return fmt.Errorf("format tanggal tidak valid: %v", err)
		}
	}

	// Menangani Waktu Input (hanya jam dan menit)
	if aux.WaktuInput != "" {
		// Menambahkan tanggal default (hari ini)
		currentDate := time.Now().Format("2006-01-02")
		waktuStr := currentDate + " " + aux.WaktuInput // Gabungkan tanggal dengan waktu
		var err error
		parsedTime, err := time.Parse("2006-01-02 15:04", waktuStr) // Format "YYYY-MM-DD HH:MM"
		if err != nil {
			return fmt.Errorf("format waktu tidak valid: %v", err)
		}
		k.Waktu = &parsedTime // Assign pointer to k.Waktu
	}

	// Tambahkan validasi jika tahun Waktu tidak valid
	if k.Waktu != nil && (k.Waktu.Year() < 1 || k.Waktu.Year() > 9999) {
		return fmt.Errorf("tahun tidak valid di waktu: %v", k.Waktu.Year())
	}

	return nil
}

package models

import (
	"time"
)

type Kendaraan struct {
	IdKendaraan   uint       `json:"idKendaraan" gorm:"primaryKey;autoIncrement;column:idKendaraan;not null"`
	NamaKendaraan string     `json:"namaKendaraan" gorm:"column:namaKendaraan;not null"`
	Warna         string     `json:"warna" gorm:"column:warna;not null"`
	Tipe          string     `json:"tipe" gorm:"column:tipe;not null"`
	PlatNomor     string     `json:"platNomor" gorm:"column:platNomor;not null"`
	DeletedAt     *time.Time `json:"deletedAt" gorm:"column:deletedAt"`
}

func (Kendaraan) TableName() string {
	return "kendaraan" // Explicitly set the table name to "lokasi"
}

package models

import (
	"time"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Mahasiswa struct {
	ID string `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null"`
	NIM int `gorm:"type:int(10);not null;unique"`
	Prodi string `gorm:"type:enum('IF', 'TI', 'TE', 'SI', 'RPL')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Mahasiswa) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()

	return
}
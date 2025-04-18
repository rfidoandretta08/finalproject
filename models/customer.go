package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"` // FK ke User
	Nama      string         `gorm:"type:varchar(255);not null" json:"nama"`
	Email     string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Alamat    string         `gorm:"type:text" json:"alamat"`
	Gender    string         `gorm:"type:enum('L','P')" json:"gender"`
	Usia      int            `json:"usia"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	User   *User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Orders []Order `gorm:"foreignKey:CustomerID" json:"orders,omitempty"`
}

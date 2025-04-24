package models

import "time"

type Product struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	NamaProduct string     `gorm:"type:varchar(255);not null" json:"nama_product"`
	CategoryID  int        `gorm:"not null" json:"category_id"`
	Harga       float64    `gorm:"type:decimal(10,2)" json:"harga"`
	Deskripsi   string     `gorm:"type:text" json:"deskripsi"`
	ImageURL    string     `gorm:"type:varchar(255)" json:"image_url"`
	Stock       int        `json:"stock"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	Category     Category      `json:"category"`
	OrderDetails []OrderDetail `gorm:"foreignKey:ProductID" json:"order_details"`
}

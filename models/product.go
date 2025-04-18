package models

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	NamaProduct string  `gorm:"type:varchar(255);not null" json:"nama_product"`
	CategoryID  int     `gorm:"not null" json:"category_id"`
	Harga       float64 `gorm:"type:decimal(10,2)" json:"harga"`
	Deskripsi   string  `gorm:"type:text" json:"deskripsi"`
	ImageURL    string  `gorm:"type:varchar(255)" json:"image_url"`
	Stock       int     `json:"stock"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`

	Category     Category      `json:"category"`
	OrderDetails []OrderDetail `gorm:"foreignKey:ProductID" json:"order_details"`
}

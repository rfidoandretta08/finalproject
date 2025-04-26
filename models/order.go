package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	CustomerID      uint           `gorm:"not null" json:"customer_id"`
	TanggalOrder    time.Time      `json:"tanggal_order"`
	Status          string         `gorm:"type:enum('diproses','dikirim','selesai');default:'diproses'" json:"status"`
	TotalHarga      float64        `gorm:"type:decimal(10,2)" json:"total_harga"`
	PaymentMethod   string         `gorm:"type:enum('Cash on Delivery','Transfer','Qris');default:'Cash on Delivery'" json:"payment_method"`
	StatusOfPayment string         `gorm:"type:enum('Menunggu Pembayaran','Sudah Dibayar', Dibatalkan);default:'Menunggu Pembayaran'" json:"status_of_payment"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Customer     Customer      `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"customer,omitempty"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID" json:"order_details,omitempty"`
}

type OrderDetail struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrderID   uint           `gorm:"not null" json:"order_id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Jumlah    int            `json:"jumlah"`
	SubTotal  float64        `gorm:"type:decimal(10,2)" json:"sub_total"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Order   Order   `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order,omitempty"`
	Product Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
}

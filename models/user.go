package models

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"type:varchar(100);not null" json:"username"`
	Password  string     `gorm:"type:varchar(255);not null" json:"password"`
	Email     string     `gorm:"type:varchar(255);not null" json:"email"`
	Phone     string     `gorm:"type:varchar(20)" json:"phone"`
	Role      string     `gorm:"type:enum('admin','customer');default:'customer'" json:"role"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	Customer []Customer `gorm:"foreignKey:UserID" json:"customer"`
}

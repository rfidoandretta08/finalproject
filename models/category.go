package models

import (
	"time"
)

type Category struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(255);not null;unique" json:"name"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}

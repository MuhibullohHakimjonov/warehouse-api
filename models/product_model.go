package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Type    string `gorm:"not null"`
	ShelfID uint
	Shelf   Shelf `gorm:"foreignKey:ShelfID"`
}

package models

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
)

type LtreePath string

func (p *LtreePath) Scan(value interface{}) error {
	*p = LtreePath(fmt.Sprintf("%s", value))
	return nil
}

func (p LtreePath) Value() (driver.Value, error) {
	return string(p), nil
}

type Shelf struct {
	gorm.Model
	ShelfParent string `gorm:"not null"`
	ShelfChild  string `gorm:"type:ltree;not null;unique"`
}

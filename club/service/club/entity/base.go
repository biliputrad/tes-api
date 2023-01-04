package entity

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}

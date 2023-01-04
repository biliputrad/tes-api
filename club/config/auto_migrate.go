package config

import (
	"gorm.io/gorm"
	"test-api/club/service/club/entity"
)

type Entity struct {
	EntityClub interface{}
}

func RegisterEntityClub() []Entity {
	return []Entity{
		{EntityClub: entity.Club{}},
	}
}

func Club(db *gorm.DB) error {
	for _, club := range RegisterEntityClub() {
		dbMigErr := db.Debug().AutoMigrate(club.EntityClub)
		if dbMigErr != nil {
			return dbMigErr
		}
	}
	return nil
}

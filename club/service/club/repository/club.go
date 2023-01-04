package repository

import (
	"gorm.io/gorm"
	"test-api/club/service/club/entity"
)

type ClubRepository interface {
	CreateClubRepository(club entity.Club) (err error)
	GetAllClubRepository() (response []entity.Club, err error)
	UpdateClubScore(clubName string, score int64) (err error)
	GetClubByName(name string) (response entity.Club, err error)
	UpdateClub(club entity.Club) (err error)
}

type clubRepository struct {
	db *gorm.DB
}

func NewClubRepository(db *gorm.DB) *clubRepository {
	return &clubRepository{db}
}

func (r *clubRepository) CreateClubRepository(club entity.Club) (err error) {
	err = r.db.Create(&club).Error

	return err
}

func (r *clubRepository) GetAllClubRepository() (response []entity.Club, err error) {
	err = r.db.Find(&response).Error

	return response, err
}

func (r *clubRepository) UpdateClubScore(clubName string, score int64) (err error) {
	err = r.db.Model(&entity.Club{}).Where("club_name = ?", clubName).Update("point", score).Error

	return err
}

func (r *clubRepository) GetClubByName(name string) (response entity.Club, err error) {
	err = r.db.Where("club_name", name).Find(&response).Error

	return response, err
}

func (r *clubRepository) UpdateClub(club entity.Club) (err error) {
	err = r.db.Save(&club).Error

	return err
}

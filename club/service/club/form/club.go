package form

import (
	"test-api/club/service/club/entity"
)

type Club struct {
	ClubName string `json:"club_name" binding:"required"`
}

type RecordMatch struct {
	ClubHomeName string `json:"club_home_name" binding:"required"`
	ClubAwayName string `json:"club_away_name" binding:"required"`
	Score        string `json:"score" binding:"required"`
}

type ContainLetter struct {
	FirstWord  string `json:"first_word" binding:"required"`
	SecondWord string `json:"second_word" binding:"required"`
}

func ConvertIntoEntityClub(club Club) (response entity.Club) {
	response.ClubName = club.ClubName

	return response
}

package response

import "test-api/club/service/club/entity"

type Club struct {
	ClubName string `json:"club_name"`
	Point    int64  `json:"point"`
	Rank     int    `json:"rank"`
}

func ConvertEntityToResponseClub(clubs []entity.Club) (response []Club) {
	for _, club := range clubs {
		responseClub := Club{
			ClubName: club.ClubName,
			Point:    club.Point,
			Rank:     club.Rank,
		}
		response = append(response, responseClub)
	}

	return response
}

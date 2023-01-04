package usecase

import (
	"regexp"
	"sort"
	"strconv"
	"test-api/club/service/club/entity"
	"test-api/club/service/club/form"
	"test-api/club/service/club/repository"
	"test-api/club/service/club/response"
)

type ClubUseCase interface {
	CreateClubUseCase(input form.Club) (result bool, err error)
	UpdateRecordMatchUserCase(input form.RecordMatch) (result bool, err error)
	updateHomeAway(match []string, home, away entity.Club) (err error)
	GetAllClubUseCase() (result []response.Club, err error)
	ContainLetterUseCase(letter form.ContainLetter) (result bool)
}

type clubUseCase struct {
	clubRepository repository.ClubRepository
}

func NewClubUseCase(clubRepository repository.ClubRepository) *clubUseCase {
	return &clubUseCase{clubRepository}
}

func (s *clubUseCase) CreateClubUseCase(input form.Club) (result bool, err error) {
	club := form.ConvertIntoEntityClub(input)

	clubs, err := s.clubRepository.GetAllClubRepository()
	if err != nil {
		return false, err
	}
	if len(clubs) == 0 {
		club.Rank = 1
	} else if len(clubs) > 1 {
		club.Rank = len(clubs) + 1
	}

	err = s.clubRepository.CreateClubRepository(club)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *clubUseCase) UpdateRecordMatchUserCase(input form.RecordMatch) (result bool, err error) {
	regex := regexp.MustCompile(`([0-9]+)\s*:\s*([0-9]+)`)
	match := regex.FindStringSubmatch(input.Score)

	home, err := s.clubRepository.GetClubByName(input.ClubHomeName)
	if err != nil {
		return false, err
	}

	away, err := s.clubRepository.GetClubByName(input.ClubAwayName)
	if err != nil {
		return false, err
	}

	err = s.updateHomeAway(match, home, away)
	if err != nil {
		return false, err
	}

	clubs, err := s.clubRepository.GetAllClubRepository()
	if err != nil {
		return false, err
	}

	sort.SliceStable(clubs, func(i, j int) bool {
		return clubs[i].Point > clubs[j].Point
	})

	for i, club := range clubs {
		club.Rank = i + 1
		err := s.clubRepository.UpdateClub(club)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *clubUseCase) updateHomeAway(match []string, home, away entity.Club) (err error) {
	scoreHome, _ := strconv.Atoi(match[1])
	scoreAway, _ := strconv.Atoi(match[2])
	if scoreHome == scoreAway {
		err = s.clubRepository.UpdateClubScore(home.ClubName, home.Point+1)
		if err != nil {
			return err
		}
		err = s.clubRepository.UpdateClubScore(away.ClubName, away.Point+1)
		if err != nil {
			return err
		}
	} else if scoreHome > scoreAway {
		err = s.clubRepository.UpdateClubScore(home.ClubName, home.Point+3)
		if err != nil {
			return err
		}
	} else if scoreHome < scoreAway {
		err = s.clubRepository.UpdateClubScore(away.ClubName, away.Point+3)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *clubUseCase) GetAllClubUseCase() (result []response.Club, err error) {
	clubs, err := s.clubRepository.GetAllClubRepository()

	sort.SliceStable(clubs, func(i, j int) bool {
		return clubs[i].Rank < clubs[j].Rank
	})

	result = response.ConvertEntityToResponseClub(clubs)
	if err != nil {
		return result, err
	}

	return result, err

}

func (s *clubUseCase) ContainLetterUseCase(letter form.ContainLetter) (result bool) {
	letterCounts := make(map[rune]int)
	for _, letter := range letter.FirstWord {
		letterCounts[letter]++
	}

	for _, letter := range letter.SecondWord {
		letterCounts[letter]--
		if letterCounts[letter] < 0 {
			return false
		}
	}

	return true
}

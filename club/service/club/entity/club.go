package entity

type Club struct {
	Base
	ClubName string `gorm:"unique"`
	Point    int64
	Rank     int
}

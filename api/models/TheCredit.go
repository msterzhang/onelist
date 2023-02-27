package models

// 电影演员及制作团队人员
type TheCredit struct {
	ID         int        `json:"id" gorm:"not null;unique"`
	TheTvID    uint       `json:"the_tv_id"`
	TheMovieID uint       `json:"the_movie_id"`
	Cast       []CastItem `json:"cast"`
	Crew       []CrewItem `json:"crew"`
}

// 主要演员们
type CastItem struct {
	ID                 int     `json:"id" gorm:"not null;unique"`
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CreditID           string  `json:"credit_id"`
	CastID             int     `json:"cast_id"`
	Character          string  `json:"character"`
	Order              int     `json:"order"`
	TheCreditID        uint    `json:"the_credits_id"`
}

// 制作团队人员
type CrewItem struct {
	ID                 int     `json:"id" gorm:"not null;unique"`
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CreditID           string  `json:"credit_id"`
	Department         string  `json:"department"`
	Job                string  `json:"job"`
	TheCreditID        uint    `json:"the_credits_id"`
}

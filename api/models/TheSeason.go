package models

type Episode struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int     `json:"id" gorm:"not null;unique"`
	Name           string  `json:"name"`
	Url            string  `json:"url"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
	TheSeasonID    uint    `json:"the_season_id"`
}

type TheSeason struct {
	IDS          string    `json:"_id"`
	AirDate      string    `json:"air_date"`
	Episodes     []Episode `json:"episodes"`
	Name         string    `json:"name"`
	Overview     string    `json:"overview"`
	ID           int       `json:"id" gorm:"not null;unique"`
	PosterPath   string    `json:"poster_path"`
	SeasonNumber int       `json:"season_number"`
	TheTvID      uint      `json:"the_tv_id"`
}

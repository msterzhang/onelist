package models

import (
	"time"

	"gorm.io/gorm"
)

type Season struct {
	ID           int    `json:"id" gorm:"not null;unique"`
	AirDate      string `json:"air_date"`
	EpisodeCount int    `json:"episode_count"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
	TheTvID      uint   `json:"the_tv_id"`
}

type LastEpisodeToAir struct {
	ID             int     `json:"id" gorm:"not null;unique"`
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
	TheTvID        uint    `json:"the_tv_id"`
}

type NextEpisodeToAir struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int     `json:"id" gorm:"not null;unique"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
	TheTvID        uint    `json:"the_tv_id"`
}

type Networks struct {
	ID            int     `json:"id" gorm:"not null;unique"`
	Name          string  `json:"name"`
	LogoPath      string  `json:"logo_path"`
	OriginCountry string  `json:"origin_country"`
	TheTvs        []TheTv `json:"the_tvs" gorm:"many2many:thetv_Networks;"`
}

type TheTv struct {
	ID                  int                  `json:"id" gorm:"not null;unique"`
	GalleryUid          string               `json:"gallery_uid"`
	Adult               bool                 `json:"adult"`
	BackdropPath        string               `json:"backdrop_path"`
	FirstAirDate        string               `json:"first_air_date"`
	Genres              []Genre              `json:"genres" gorm:"many2many:thetv_Genres;"`
	Homepage            string               `json:"homepage"`
	InProduction        bool                 `json:"in_production"`
	LastAirDate         string               `json:"last_air_date"`
	LastEpisodeToAir    LastEpisodeToAir     `json:"last_episode_to_air"`
	Name                string               `json:"name"`
	NextEpisodeToAir    NextEpisodeToAir     `json:"next_episode_to_air"`
	Networks            []Networks           `json:"networks"  gorm:"many2many:thetv_Networks;"`
	NumberOfEpisodes    int                  `json:"number_of_episodes"`
	NumberOfSeasons     int                  `json:"number_of_seasons"`
	OriginalLanguage    string               `json:"original_language"`
	OriginalName        string               `json:"original_name"`
	Overview            string               `json:"overview"`
	Popularity          float64              `json:"popularity"`
	PosterPath          string               `json:"poster_path"`
	ProductionCompanies []ProductionCompanie `json:"production_companies" gorm:"many2many:thetv_ProductionCompanies;"`
	ProductionCountries []ProductionCountrie `json:"production_countries" gorm:"many2many:thetv_ProductionCountries;"`
	Seasons             []Season             `json:"seasons"`
	SpokenLanguages     []SpokenLanguage     `json:"spoken_languages" gorm:"many2many:thetv_SpokenLanguages;"`
	ThePersons          []ThePerson          `json:"the_persons" gorm:"many2many:thetv_ThePersons;"`
	TheSeasons          []TheSeason          `json:"the_seasons"`
	Status              string               `json:"status"`
	Tagline             string               `json:"tagline"`
	Type                string               `json:"type"`
	VoteAverage         float64              `json:"vote_average"`
	VoteCount           int                  `json:"vote_count"`
	TheCredit           TheCredit            `json:"the_credit"`
	Star                bool                 `json:"star"`
	Heart               bool                 `json:"heart"`
	Played              bool                 `json:"played"`
	CreatedAt           time.Time            `json:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at"`
}

func (t *TheTv) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return
}

func (t *TheTv) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return
}

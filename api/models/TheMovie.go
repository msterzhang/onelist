package models

import (
	"time"

	"gorm.io/gorm"
)

// 电影剧集
type BelongsToCollection struct {
	ID           int        `json:"id" gorm:"not null;unique"`
	Name         string     `json:"name"`
	PosterPath   string     `json:"poster_path"`
	BackdropPath string     `json:"backdrop_path"`
	TheMovies    []TheMovie `json:"the_movies"`
}

// 电影结构体
type TheMovie struct {
	ID                    int                  `json:"id" gorm:"not null;unique"`
	GalleryUid            string               `json:"gallery_uid"`
	Adult                 bool                 `json:"adult"`
	BackdropPath          string               `json:"backdrop_path"`
	BelongsToCollection   BelongsToCollection  `json:"belongs_to_collection"`
	Budget                int                  `json:"budget"`
	Genres                []Genre              `json:"genres" gorm:"many2many:themovie_Genres;"`
	Homepage              string               `json:"homepage"`
	ImdbID                string               `json:"imdb_id"`
	OriginalLanguage      string               `json:"original_language"`
	OriginalTitle         string               `json:"original_title"`
	Overview              string               `json:"overview"`
	Popularity            float64              `json:"popularity"`
	PosterPath            string               `json:"poster_path"`
	ProductionCompanies   []ProductionCompanie `json:"production_companies" gorm:"many2many:themovie_ProductionCompanies;"`
	ProductionCountries   []ProductionCountrie `json:"production_countries" gorm:"many2many:themovie_ProductionCountries;"`
	ReleaseDate           string               `json:"release_date"`
	Revenue               int                  `json:"revenue"`
	Runtime               int                  `json:"runtime"`
	SpokenLanguages       []SpokenLanguage     `json:"spoken_languages" gorm:"many2many:themovie_SpokenLanguages;"`
	ThePersons            []ThePerson          `json:"the_persons" gorm:"many2many:themovie_ThePersons;"`
	Status                string               `json:"status"`
	Tagline               string               `json:"tagline"`
	Title                 string               `json:"title"`
	Url                   string               `json:"url"`
	Video                 bool                 `json:"video"`
	VoteAverage           float64              `json:"vote_average"`
	VoteCount             int                  `json:"vote_count"`
	TheCredit             TheCredit            `json:"the_credit"`
	BelongsToCollectionID uint                 `json:"belongs_to_collection_id"`
	Star                  bool                 `json:"star"`
	Heart                 bool                 `json:"heart"`
	Played                bool                 `json:"played"`
	CreatedAt             time.Time            `json:"created_at"`
	UpdatedAt             time.Time            `json:"updated_at"`
}

func (m *TheMovie) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *TheMovie) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

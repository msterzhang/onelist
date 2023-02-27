package models

// 演员信息
type ThePerson struct {
	ID                 int        `json:"id" gorm:"not null;unique"`
	Adult              bool       `json:"adult"`
	Biography          string     `json:"biography"`
	Birthday           string     `json:"birthday"`
	Deathday           string     `json:"deathday"`
	Gender             int        `json:"gender"`
	Homepage           string     `json:"homepage"`
	ImdbID             string     `json:"imdb_id"`
	KnownForDepartment string     `json:"known_for_department"`
	Name               string     `json:"name"`
	PlaceOfBirth       string     `json:"place_of_birth"`
	Popularity         float64    `json:"popularity"`
	ProfilePath        string     `json:"profile_path"`
	TheMovies          []TheMovie `json:"the_movies" gorm:"many2many:themovie_ThePersons;"`
	TheTvs             []TheTv    `json:"the_tvs" gorm:"many2many:thetv_ThePersons;"`
}

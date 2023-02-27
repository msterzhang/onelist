package models

// 电影类型
type Genre struct {
	ID        int        `json:"id" gorm:"not null;unique"`
	Name      string     `json:"name"`
	TheMovies []TheMovie `json:"the_movies" gorm:"many2many:themovie_Genres;"`
	TheTvs    []TheTv    `json:"the_tvs" gorm:"many2many:thetv_Genres;"`
}

// 制作公司
type ProductionCompanie struct {
	ID            int        `json:"id" gorm:"not null;unique"`
	LogoPath      string     `json:"logo_path"`
	Name          string     `json:"name"`
	OriginCountry string     `json:"origin_country"`
	TheMovies     []TheMovie `json:"the_movies" gorm:"many2many:themovie_ProductionCompanies;"`
	TheTvs        []TheTv    `json:"the_tvs" gorm:"many2many:thetv_ProductionCompanies;"`
}

// 制作国家
type ProductionCountrie struct {
	ID        int        `json:"id" gorm:"not null;unique"`
	Iso31661  string     `json:"iso_3166_1"`
	Name      string     `json:"name"`
	TheMovies []TheMovie `json:"the_movies" gorm:"many2many:themovie_ProductionCountries;"`
	TheTvs    []TheTv    `json:"the_tvs" gorm:"many2many:thetv_ProductionCountries;"`
}

// 发布语言
type SpokenLanguage struct {
	ID          int        `json:"id" gorm:"not null;unique"`
	EnglishName string     `json:"english_name"`
	Iso6391     string     `json:"iso_639_1"`
	Name        string     `json:"name"`
	TheMovies   []TheMovie `json:"the_movies" gorm:"many2many:themovie_SpokenLanguages;"`
	TheTvs      []TheTv    `json:"the_tvs" gorm:"many2many:thetv_SpokenLanguages;"`
}

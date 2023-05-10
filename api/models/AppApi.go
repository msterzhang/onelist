package models

type TheDataIndex struct {
	Title        string     `json:"title"`
	GalleryType  string     `json:"gallery_type"`
	TheMovieList []TheMovie `json:"the_movie_list"`
	TheTvList    []TheTv    `json:"the_tv_list"`
}

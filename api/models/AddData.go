package models

type AddVideo struct {
	TheMovieId int    `json:"the_movie_id"`
	TheTvId    int    `json:"the_tv_id"`
	GalleryUid string `json:"gallery_uid"`
	Path       string `json:"path"`
	File       string `json:"file"`
}

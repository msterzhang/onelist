package service

import (
	"errors"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"

	"gorm.io/gorm"
)

// 处理电视用户是否已点赞，收藏及点了最爱
func TheTvService(theTv models.TheTv, userId string) models.TheTv {
	star := models.Star{}
	played := models.Played{}
	heart := models.Heart{}
	db := database.NewDb()
	err := db.Model(&models.Star{}).Where("user_id = ? AND data_id = ? AND data_type = ?", userId, theTv.ID, "tv").First(&star).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && star.Id != 0 {
		theTv.Star = true
	}
	err = db.Model(&models.Played{}).Where("user_id = ? AND data_id = ? AND data_type = ?", userId, theTv.ID, "tv").First(&played).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && played.Id != 0 {
		theTv.Played = true
	}
	err = db.Model(&models.Heart{}).Where("user_id = ? AND data_id = ? AND data_type = ?", userId, theTv.ID, "tv").First(&heart).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && heart.Id != 0 {
		theTv.Heart = true
	}
	return theTv
}

// 处理电视用户是否已点赞，收藏及点了最爱
func TheTvsService(theTvs []models.TheTv, userId string) []models.TheTv {
	var newTheTvs []models.TheTv
	for _, theTv := range theTvs {
		newTheTvs = append(newTheTvs, TheTvService(theTv, userId))
	}
	return newTheTvs
}

// 处理电影用户是否已点赞，收藏及点了最爱
func TheMovieService(theMovie models.TheMovie, userId string) models.TheMovie {
	star := models.Star{}
	played := models.Played{}
	heart := models.Heart{}
	db := database.NewDb()
	err := db.Model(&models.Star{}).Where("user_id = ? AND data_id = ? AND data_type = ?", userId, theMovie.ID, "movie").First(&star).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && star.Id != 0 {
		theMovie.Star = true
	}
	err = db.Model(&models.Played{}).Where("user_id = ? AND data_id = ? AND data_type = ?", userId, theMovie.ID, "movie").First(&played).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && played.Id != 0 {
		theMovie.Played = true
	}
	err = db.Model(&models.Heart{}).Where("user_id = ? AND data_id = ? AND data_type = ?", userId, theMovie.ID, "movie").First(&heart).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && heart.Id != 0 {
		theMovie.Heart = true
	}
	return theMovie
}

// 处理电影用户是否已点赞，收藏及点了最爱
func TheMoviesService(theMovies []models.TheMovie, userId string) []models.TheMovie {
	var newTheMovies []models.TheMovie
	for _, theMovie := range theMovies {
		newTheMovies = append(newTheMovies, TheMovieService(theMovie, userId))
	}
	return newTheMovies
}

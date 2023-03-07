package thedb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/utils/extract"
	"github.com/msterzhang/onelist/config"

	"gorm.io/gorm"
)

// 99.84.251.12 api.themoviedb.org
// 99.84.251.19 api.themoviedb.org
// 99.84.251.67 api.themoviedb.org
// 99.84.251.108 api.themoviedb.org
// 156.146.56.162 image.tmdb.org
// 108.138.246.49 image.tmdb.org
// https://image.tmdb.org/t/p/w220_and_h330_face/h7thH2xZeicwK7a3Pkr4cCzXkSu.jpg
// https://image.tmdb.org/t/p/w1920_and_h1080_bestv2/yL0h5NggYqBzGvLzc4TTM049jDm.jpg
// https://image.tmdb.org/t/p/w355_and_h200_multi_faces/yL0h5NggYqBzGvLzc4TTM049jDm.jpg
// https://image.tmdb.org/t/p/w227_and_h127_bestv2/i5LwCRuHRuQxPVJPbOAIIkRKiQo.jpg

var (
	ImageHost = "http://image.tmdb.org/"
	TheApi    = "https://api.themoviedb.org/3"
	// 取0-24，共计24人
	personNumber = 24
	timeOut      = 30 * time.Second
)

// 搜索资源
func SearchTheDb(key string, tv bool) (ThedbSearchRsp, error) {
	if !tv {
		key = extract.ExtractMovieName(key)
	}
	api := fmt.Sprintf("%s/search/movie?api_key=%s&language=zh&page=1&query=%s", TheApi, config.KeyDb, key)
	if tv {
		api = fmt.Sprintf("%s/search/tv?api_key=%s&language=zh&page=1&query=%s", TheApi, config.KeyDb, key)
	}
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return ThedbSearchRsp{}, err
	}
	req.Header.Set("User-Agent", config.UA)
	client := http.Client{
		Timeout: timeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		return ThedbSearchRsp{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ThedbSearchRsp{}, err
	}
	var data = ThedbSearchRsp{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return ThedbSearchRsp{}, err
	}
	return data, nil
}

// 获取整个剧组人员
func GetCredits(id int, tv bool) (models.TheCredit, error) {
	api := fmt.Sprintf("%s/movie/%d/credits?api_key=%s&language=zh", TheApi, id, config.KeyDb)
	if tv {
		api = fmt.Sprintf("%s/tv/%d/credits?api_key=%s&language=zh", TheApi, id, config.KeyDb)
	}
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return models.TheCredit{}, err
	}
	req.Header.Set("User-Agent", config.UA)
	client := http.Client{
		Timeout: timeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		return models.TheCredit{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.TheCredit{}, err
	}
	var data = models.TheCredit{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return models.TheCredit{}, err
	}
	cast := personNumber
	crew := personNumber
	if len(data.Cast) <= cast {
		cast = len(data.Cast)
	}
	if len(data.Crew) <= crew {
		crew = len(data.Crew)
	}
	data.Cast = data.Cast[:cast]
	data.Crew = data.Crew[:crew]
	return data, nil
}

// 获取电影数据
func GetMovieData(id int) (models.TheMovie, error) {
	api := fmt.Sprintf("%s/movie/%d?api_key=%s&language=zh", TheApi, id, config.KeyDb)
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return models.TheMovie{}, err
	}
	req.Header.Set("User-Agent", config.UA)
	client := http.Client{
		Timeout: timeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		return models.TheMovie{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.TheMovie{}, err
	}
	var data = models.TheMovie{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return models.TheMovie{}, err
	}
	if config.DownLoadImage == "是" {
		DownImages(data.PosterPath)
		DownBackImage(data.BackdropPath)
	}
	return data, nil
}

// 获取电视节目数据
func GetTvData(id int) (models.TheTv, error) {
	api := fmt.Sprintf("%s/tv/%d?api_key=%s&language=zh", TheApi, id, config.KeyDb)
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return models.TheTv{}, err
	}
	req.Header.Set("User-Agent", config.UA)
	client := http.Client{
		Timeout: timeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		return models.TheTv{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.TheTv{}, err
	}
	var data = models.TheTv{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return models.TheTv{}, err
	}
	if config.DownLoadImage == "是" {
		DownImages(data.PosterPath)
		DownBackImage(data.BackdropPath)
	}
	return data, nil
}

// 获取电视每季详情
func GetTheSeasonData(id int, item int) (models.TheSeason, error) {
	api := fmt.Sprintf("%s/tv/%d/season/%d?api_key=%s&language=zh", TheApi, id, item, config.KeyDb)
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return models.TheSeason{}, err
	}
	req.Header.Set("User-Agent", config.UA)
	client := http.Client{
		Timeout: timeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		return models.TheSeason{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.TheSeason{}, err
	}
	if strings.Contains(string(body), "not be found") {
		return models.TheSeason{}, errors.New("not be found")
	}
	var data = models.TheSeason{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return models.TheSeason{}, err
	}
	if config.DownLoadImage == "是" {
		DownSeasonImages(data.PosterPath)
	}
	return data, nil
}

// 获取演员信息
func GetThePersonData(id int) (models.ThePerson, error) {
	api := fmt.Sprintf("%s/person/%d?api_key=%s&language=zh", TheApi, id, config.KeyDb)
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return models.ThePerson{}, err
	}
	req.Header.Set("User-Agent", config.UA)
	client := http.Client{
		Timeout: timeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		return models.ThePerson{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ThePerson{}, err
	}
	var data = models.ThePerson{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return models.ThePerson{}, err
	}
	if config.DownLoadImage == "是" {
		DownPersonImage(data.ProfilePath)
	}
	return data, nil
}

// 检查是否已存在此人，存在则更新，不存在则创建,注意人物和电影及电视都存在着关联
func ChunkPerson(person models.ThePerson) error {
	db := database.NewDb()
	dbPerson := models.ThePerson{}
	err := db.Model(&models.ThePerson{}).Where("id = ?", person.ID).First(&dbPerson).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Model(&models.ThePerson{}).Create(&person).Error
	}
	err = db.Model(&models.ThePerson{}).Where("id = ?", person.ID).Select("*").Updates(&person).Error
	return err
}

// 检查是否已存在此电影，存在则更新，不存在则创建
func ChunkTheMovie(themovie models.TheMovie) error {
	db := database.NewDb()
	dbThemovie := models.TheMovie{}
	err := db.Model(&models.TheMovie{}).Where("id = ?", themovie.ID).First(&dbThemovie).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Debug().Model(&models.TheMovie{}).Create(&themovie).Error
	}
	themovie.CreatedAt = dbThemovie.CreatedAt
	err = db.Model(&models.TheMovie{}).Where("id = ?", themovie.ID).Select("*").Updates(&themovie).Error
	return err
}

// 根据电影ID及文件刮削保存资源
func TheMovieDb(id int, file string, GalleryUid string) (models.TheMovie, error) {
	data, err := GetMovieData(id)
	if err != nil {
		return models.TheMovie{}, err
	}
	credit, err := GetCredits(id, false)
	if err != nil {
		return models.TheMovie{}, err
	}
	data.TheCredit = credit
	casts := credit.Cast
	// persons := []models.ThePerson{}
	db := database.NewDb()
	for _, cast := range casts {
		dbPerson := models.ThePerson{}
		err := db.Model(&models.ThePerson{}).Where("id = ?", cast.ID).First(&dbPerson).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			porson, err := GetThePersonData(cast.ID)
			if err != nil {
				continue
			}
			porson.TheMovies = append(porson.TheMovies, data)
			err = ChunkPerson(porson)
			if err != nil {
				continue
			}
		}
	}
	crews := credit.Crew
	for _, crew := range crews {
		dbPerson := models.ThePerson{}
		err := db.Model(&models.ThePerson{}).Where("id = ?", crew.ID).First(&dbPerson).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			porson, err := GetThePersonData(crew.ID)
			if err != nil {
				continue
			}
			porson.TheMovies = append(porson.TheMovies, data)
			err = ChunkPerson(porson)
			if err != nil {
				continue
			}
		}
	}
	data.Url = file
	data.GalleryUid = GalleryUid
	err = ChunkTheMovie(data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// 根据节目数据获取指定季的信息
func GetSeasonWithTheTv(theTv models.TheTv, item int) (models.Season, error) {
	for _, line := range theTv.Seasons {
		if line.SeasonNumber == item {
			return line, nil
		}
	}
	return models.Season{}, errors.New("not find")
}

// 根据每一季数据获取指定集的信息
func GetEpisodeWithTheSeason(season models.TheSeason, item int) (models.Episode, error) {
	for _, line := range season.Episodes {
		if line.EpisodeNumber == item {
			return line, nil
		}
	}
	return models.Episode{}, errors.New("not find")
}

// 根据电视剧Id及文件刮削保存资源
func TheTvDb(id int, file string, GalleryUid string) (models.TheTv, error) {
	data, err := GetTvData(id)
	if err != nil {
		return models.TheTv{}, err
	}
	credit, err := GetCredits(id, true)
	if err != nil {
		return models.TheTv{}, err
	}
	data.TheCredit = credit
	casts := credit.Cast
	// persons := []models.ThePerson{}
	for _, cast := range casts {
		porson, err := GetThePersonData(cast.ID)
		if err != nil {
			continue
		}
		porson.TheTvs = append(porson.TheTvs, data)
		err = ChunkPerson(porson)
		if err != nil {
			continue
		}
	}
	crews := credit.Crew
	for _, crew := range crews {
		porson, err := GetThePersonData(crew.ID)
		if err != nil {
			continue
		}
		porson.TheTvs = append(porson.TheTvs, data)
		err = ChunkPerson(porson)
		if err != nil {
			continue
		}
	}
	SeasonNumber, EpisodeNumber, err := extract.ExtractNumberWithFile(file)
	if err != nil {
		return models.TheTv{}, err
	}
	theseason, err := GetTheSeasonData(id, SeasonNumber)
	if err != nil {
		return models.TheTv{}, err
	}
	season, err := GetSeasonWithTheTv(data, SeasonNumber)
	if err != nil {
		return models.TheTv{}, err
	}
	episode, err := GetEpisodeWithTheSeason(theseason, EpisodeNumber)
	if err != nil {
		return models.TheTv{}, err
	}
	DownEpisodeImages(episode.StillPath)
	season.TheTvID = uint(data.ID)
	theseason.TheTvID = uint(data.ID)
	episode.TheSeasonID = uint(theseason.ID)
	episode.Url = file
	data.Seasons = []models.Season{}
	theseason.Episodes = []models.Episode{}
	err = ChunkSeason(season)
	if err != nil {
		return data, err
	}
	err = ChunkTheSeason(theseason)
	if err != nil {
		return data, err
	}
	err = ChunkEpisode(episode)
	if err != nil {
		return data, err
	}
	data.GalleryUid = GalleryUid
	err = ChunkTheTv(data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// 检查是否已存在此节目，存在则更新，不存在则创建
func ChunkTheTv(thetv models.TheTv) error {
	db := database.NewDb()
	dbthetv := models.TheTv{}
	err := db.Model(&models.TheTv{}).Where("id = ?", thetv.ID).First(&dbthetv).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Model(&models.TheTv{}).Create(&thetv).Error
	}
	thetv.CreatedAt = dbthetv.CreatedAt
	err = db.Model(&models.TheTv{}).Where("id = ?", thetv.ID).Select("*").Updates(&thetv).Error
	return err
}

// 检查是否已存在此节目分季，存在则更新，不存在则创建
func ChunkTheSeason(theseason models.TheSeason) error {
	db := database.NewDb()
	dbtheseason := models.TheSeason{}
	err := db.Model(&models.TheSeason{}).Where("id = ?", theseason.ID).First(&dbtheseason).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Model(&models.TheSeason{}).Create(&theseason).Error
	}
	err = db.Model(&models.TheSeason{}).Where("id = ?", theseason.ID).Select("*").Updates(&theseason).Error
	return err
}

// 检查是否已存在此节目分集，存在则更新，不存在则创建
func ChunkEpisode(episode models.Episode) error {
	db := database.NewDb()
	dbepisode := models.Episode{}
	err := db.Model(&models.Episode{}).Where("id = ?", episode.ID).First(&dbepisode).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Model(&models.Episode{}).Create(&episode).Error
	}
	err = db.Model(&models.Episode{}).Where("id = ?", episode.ID).Select("*").Updates(&episode).Error
	return err
}

// 检查是否已存在此节目分季，存在则更新，不存在则创建
func ChunkSeason(season models.Season) error {
	db := database.NewDb()
	dbseason := models.Season{}
	err := db.Model(&models.Season{}).Where("id = ?", season.ID).First(&dbseason).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Model(&models.Season{}).Create(&season).Error
	}
	err = db.Model(&models.Season{}).Where("id = ?", season.ID).Select("*").Updates(&season).Error
	return err
}

// 自动刮削保存电影
func RunTheMovieWork(file string, GalleryUid string) (int, error) {
	p, err := filepath.Abs(file)
	if err != nil {
		return 0, err
	}
	fileName := filepath.Base(p)
	fileType := path.Ext(fileName)
	name := strings.ReplaceAll(fileName, fileType, "")
	data, err := SearchTheDb(name, false)
	if err != nil {
		return 0, err
	}
	if len(data.Results) == 0 {
		return 0, errors.New("movie not found")
	}
	id := data.Results[0].ID
	_, err = TheMovieDb(id, file, GalleryUid)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 自动刮削保存电视
func RunTheTvWork(file string, GalleryUid string) (int, error) {
	p, err := filepath.Abs(file)
	if err != nil {
		return 0, err
	}
	fileName := filepath.Base(p)
	fileType := path.Ext(fileName)
	name := strings.ReplaceAll(fileName, fileType, "")
	re := regexp.MustCompile(`[\p{Han}\d{1,2}]+`)
	matches := re.FindAllString(name, -1)
	if len(matches) > 0 {
		name = matches[0]
	}
	data, err := SearchTheDb(name, true)
	if err != nil {
		return 0, err
	}
	if len(data.Results) == 0 {
		return 0, errors.New("tv not found")
	}
	id := data.Results[0].ID
	thetv, err := TheTvDb(id, file, GalleryUid)
	if err != nil {
		return 0, err
	}
	return thetv.ID, nil
}

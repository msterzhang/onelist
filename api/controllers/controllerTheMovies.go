package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"
	"github.com/msterzhang/onelist/api/service"
	"github.com/msterzhang/onelist/plugins/thedb"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTheMovie(c *gin.Context) {
	themovie := models.TheMovie{}
	err := c.ShouldBind(&themovie)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": themovie})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheMoviesCRUD(db)
	func(themovieRepository repository.TheMovieRepository) {
		themovie, err := themovieRepository.Save(themovie)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": themovie})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": themovie})
	}(repo)
}

func DeleteTheMovieById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryTheMoviesCRUD(db)
	func(themovieRepository repository.TheMovieRepository) {
		themovie, err := themovieRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": themovie})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": themovie})
	}(repo)
}

func UpdateTheMovieById(c *gin.Context) {
	id := c.Query("id")
	themovie := models.TheMovie{}
	err := c.ShouldBind(&themovie)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": themovie})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheMoviesCRUD(db)
	func(themovieRepository repository.TheMovieRepository) {
		themovie, err := themovieRepository.UpdateByID(id, themovie)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": themovie})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": themovie})
	}(repo)
}

func GetTheMovieById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryTheMoviesCRUD(db)
	func(themovieRepository repository.TheMovieRepository) {
		themovie, err := themovieRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": themovie})
			return
		}
		themovieNew := service.TheMovieService(themovie, c.GetString("UserId"))
		tag := "剧情"
		if len(themovieNew.Genres) > 1 {
			if themovieNew.Genres[0].Name == tag {
				tag = themovieNew.Genres[1].Name
			} else {
				tag = themovieNew.Genres[0].Name
			}
		}
		genre := models.Genre{}
		db.Where("name = ?", tag).Preload("TheMovies", func(db *gorm.DB) *gorm.DB {
			return db.Order("datetime(updated_at) desc").Limit(12)
		}).Find(&genre)
		genre.TheMovies = service.TheMoviesService(genre.TheMovies, c.GetString("UserId"))
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": themovieNew, "like": genre.TheMovies})
	}(repo)
}

func GetTheMovieList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheMoviesCRUD(db)
	func(themovieRepository repository.TheMovieRepository) {
		themovies, num, err := themovieRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": themovies, "num": num})
			return
		}
		themoviesNew := service.TheMoviesService(themovies, c.GetString("UserId"))
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": themoviesNew, "num": num})
	}(repo)
}

func SearchTheMovie(c *gin.Context) {
	q := c.Query("q")
	if len(q) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误!", "data": ""})
		return
	}
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheMoviesCRUD(db)
	func(themovieRepository repository.TheMovieRepository) {
		themovies, num, err := themovieRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": themovies, "num": num})
			return
		}
		themoviesNew := service.TheMoviesService(themovies, c.GetString("UserId"))
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": themoviesNew, "num": num})
	}(repo)
}

func TheMovieFilter(c *gin.Context) {
	q := c.Query("q")
	if len(q) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误!", "data": ""})
		return
	}
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheMoviesCRUD(db)
	func(themovieRepository repository.TheMovieRepository) {
		themovies, num, err := themovieRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": themovies, "num": num})
			return
		}
		themoviesNew := service.TheMoviesService(themovies, c.GetString("UserId"))
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": themoviesNew, "num": num})
	}(repo)
}

func SortThemovie(c *gin.Context) {
	galleryUid := c.Query("gallery_uid")
	if len(galleryUid) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数不足!", "data": ""})
		return
	}
	order := c.Query("order")
	if len(order) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数不足!", "data": ""})
		return
	}
	mode := c.Query("mode")
	if len(mode) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数不足!", "data": ""})
		return
	}
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheMoviesCRUD(db)
	func(themovieRepository repository.TheMovieRepository) {
		themovies, num, err := themovieRepository.Sort(galleryUid, mode, order, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": themovies, "num": num})
			return
		}
		themoviesNew := service.TheMoviesService(themovies, c.GetString("UserId"))
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": themoviesNew, "num": num})
	}(repo)
}

func GetTheMovieListByGalleryId(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误!", "data": ""})
		return
	}
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTheMoviesCRUD(db)
	func(themovieRepository repository.TheMovieRepository) {
		themovies, num, err := themovieRepository.FindByGalleryId(id, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": themovies, "num": num})
			return
		}
		themoviesNew := service.TheMoviesService(themovies, c.GetString("UserId"))
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": themoviesNew, "num": num})
	}(repo)
}

// 手动添加视频
func AddThemovie(c *gin.Context) {
	addVideo := models.AddVideo{}
	err := c.ShouldBind(&addVideo)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "添加资源失败,表单解析出错!", "data": ""})
		return
	}
	gallery := models.Gallery{}
	db := database.NewDb()
	err = db.Model(&models.Gallery{}).Where("gallery_uid = ?", addVideo.GalleryUid).First(&gallery).Error
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "Gallery not found!", "data": ""})
		return
	}
	file := addVideo.File
	if gallery.IsAlist {
		file = "/d" + addVideo.File
	}
	themovieNew, err := thedb.TheMovieDb(addVideo.TheMovieId, file, addVideo.GalleryUid)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "添加资源失败!", "data": err})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "刮削电影成功!", "data": themovieNew.ID})
}

package api

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msterzhang/onelist/api/auth"
	"github.com/msterzhang/onelist/api/controllers"
	"github.com/msterzhang/onelist/api/crons"
	"github.com/msterzhang/onelist/api/middleware"
	"github.com/msterzhang/onelist/auto"
	"github.com/msterzhang/onelist/config"
	"github.com/msterzhang/onelist/public"
)

// 初始化配置及数据库
func InitServer() {
	config.Load()
	auto.Load()
	crons.Load()
}

// 用于打包的静态文件
func Static(r *gin.Engine) {
	folders := []string{"js", "css", "images", "fonts", "img"}
	for i, folder := range folders {
		folder = "dist/" + folder
		sub, err := fs.Sub(public.Public, folder)
		if err != nil {
			log.Fatalf("can't find folder: %s", folder)
		}
		r.StaticFS(fmt.Sprintf("/%s/", folders[i]), http.FS(sub))
	}
}

func IndexView(c *gin.Context) {
	c.Writer.WriteHeader(200)
	b, _ := public.Public.ReadFile("dist/index.html")
	_, _ = c.Writer.Write(b)
	c.Writer.Header().Add("Accept", "text/html")
	c.Writer.Flush()
}

func Faviconico(c *gin.Context) {
	c.Redirect(302, config.FaviconicoUrl)
}

func Run() {
	// 初始化
	InitServer()

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.SetMode(gin.ReleaseMode)

	//系统初始化
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	Static(r)
	r.GET("/favicon.ico", Faviconico)
	r.GET("/", IndexView)
	r.NoRoute(IndexView)

	// 用户
	user := r.Group("/v1/api/user")
	user.POST("/create", controllers.CreateUser)
	user.POST("/login", controllers.LoginUser)
	user.GET("/data", auth.JWTAuth(), controllers.UserData)
	user.POST("/update", auth.JWTAuth(), controllers.UpdateUserById)
	user.POST("/delete", auth.JWTAuthAdmin(), controllers.DeleteUserById)
	user.POST("/id", auth.JWTAuth(), controllers.GetUserById)
	user.POST("/list", auth.JWTAuthAdmin(), controllers.GetUserList)
	user.POST("/search", auth.JWTAuthAdmin(), controllers.SearchUser)

	// 标签管理
	genre := r.Group("/v1/api/genre")
	genre.POST("/create", auth.JWTAuthAdmin(), controllers.CreateGenre)
	genre.POST("/update", auth.JWTAuthAdmin(), controllers.UpdateGenreById)
	genre.POST("/delete", auth.JWTAuthAdmin(), controllers.DeleteGenreById)
	genre.POST("/id", auth.JWTAuth(), controllers.GetGenreById)
	genre.POST("/list", auth.JWTAuth(), controllers.GetGenreList)
	genre.POST("/search", auth.JWTAuth(), controllers.SearchGenre)
	genre.POST("/filte", auth.JWTAuth(), controllers.GetByIdFilte)

	productioncompanie := r.Group("/v1/api/productioncompanie", auth.JWTAuthAdmin())
	productioncompanie.POST("/create", controllers.CreateProductionCompanie)
	productioncompanie.POST("/update", controllers.UpdateProductionCompanieById)
	productioncompanie.POST("/delete", controllers.DeleteProductionCompanieById)
	productioncompanie.POST("/id", controllers.GetProductionCompanieById)
	productioncompanie.POST("/list", controllers.GetProductionCompanieList)
	productioncompanie.POST("/search", controllers.SearchProductionCompanie)

	productioncountrie := r.Group("/v1/api/productioncountrie", auth.JWTAuthAdmin())
	productioncountrie.POST("/create", controllers.CreateProductionCountrie)
	productioncountrie.POST("/update", controllers.UpdateProductionCountrieById)
	productioncountrie.POST("/delete", controllers.DeleteProductionCountrieById)
	productioncountrie.POST("/id", controllers.GetProductionCountrieById)
	productioncountrie.POST("/list", controllers.GetProductionCountrieList)
	productioncountrie.POST("/search", controllers.SearchProductionCountrie)

	// 发布语言
	spokenlanguage := r.Group("/v1/api/spokenlanguage", auth.JWTAuthAdmin())
	spokenlanguage.POST("/create", controllers.CreateSpokenLanguage)
	spokenlanguage.POST("/update", controllers.UpdateSpokenLanguageById)
	spokenlanguage.POST("/delete", controllers.DeleteSpokenLanguageById)
	spokenlanguage.POST("/id", controllers.GetSpokenLanguageById)
	spokenlanguage.POST("/list", controllers.GetSpokenLanguageList)
	spokenlanguage.POST("/search", controllers.SearchSpokenLanguage)

	// 剧组人员
	thecredit := r.Group("/v1/api/thecredit", auth.JWTAuthAdmin())
	thecredit.POST("/create", controllers.CreateTheCredit)
	thecredit.POST("/update", controllers.UpdateTheCreditById)
	thecredit.POST("/delete", controllers.DeleteTheCreditById)
	thecredit.POST("/id", controllers.GetTheCreditById)
	thecredit.POST("/list", controllers.GetTheCreditList)
	thecredit.POST("/search", controllers.SearchTheCredit)

	castitem := r.Group("/v1/api/castitem", auth.JWTAuthAdmin())
	castitem.POST("/create", controllers.CreateCastItem)
	castitem.POST("/update", controllers.UpdateCastItemById)
	castitem.POST("/delete", controllers.DeleteCastItemById)
	castitem.POST("/id", controllers.GetCastItemById)
	castitem.POST("/list", controllers.GetCastItemList)
	castitem.POST("/search", controllers.SearchCastItem)

	crewitem := r.Group("/v1/api/crewitem", auth.JWTAuthAdmin())
	crewitem.POST("/create", controllers.CreateCrewItem)
	crewitem.POST("/update", controllers.UpdateCrewItemById)
	crewitem.POST("/delete", controllers.DeleteCrewItemById)
	crewitem.POST("/id", controllers.GetCrewItemById)
	crewitem.POST("/list", controllers.GetCrewItemList)
	crewitem.POST("/search", controllers.SearchCrewItem)

	belongstocollection := r.Group("/v1/api/belongstocollection", auth.JWTAuthAdmin())
	belongstocollection.POST("/create", controllers.CreateBelongsToCollection)
	belongstocollection.POST("/update", controllers.UpdateBelongsToCollectionById)
	belongstocollection.POST("/delete", controllers.DeleteBelongsToCollectionById)
	belongstocollection.POST("/id", controllers.GetBelongsToCollectionById)
	belongstocollection.POST("/list", controllers.GetBelongsToCollectionList)
	belongstocollection.POST("/search", controllers.SearchBelongsToCollection)

	// 电影
	themovie := r.Group("/v1/api/themovie", auth.JWTAuth())
	themovie.POST("/create", controllers.CreateTheMovie)
	themovie.POST("/update", controllers.UpdateTheMovieById)
	themovie.POST("/delete", controllers.DeleteTheMovieById)
	themovie.POST("/id", controllers.GetTheMovieById)
	themovie.POST("/list", controllers.GetTheMovieList)
	themovie.POST("/gallery/list", controllers.GetTheMovieListByGalleryId)
	themovie.POST("/search", controllers.SearchTheMovie)
	themovie.POST("/add", controllers.AddThemovie)
	themovie.POST("/sort", controllers.SortThemovie)

	// 演员
	theperson := r.Group("/v1/api/theperson", auth.JWTAuth())
	theperson.POST("/create", controllers.CreateThePerson)
	theperson.POST("/update", controllers.UpdateThePersonById)
	theperson.POST("/delete", controllers.DeleteThePersonById)
	theperson.POST("/id", controllers.GetThePersonById)
	theperson.POST("/list", controllers.GetThePersonList)
	theperson.POST("/search", controllers.SearchThePerson)

	// 电视剧
	thetv := r.Group("/v1/api/thetv", auth.JWTAuth())
	thetv.POST("/create", controllers.CreateTheTv)

	thetv.POST("/update", controllers.UpdateTheTvById)
	thetv.POST("/delete", controllers.DeleteTheTvById)
	thetv.POST("/id", controllers.GetTheTvById)
	thetv.POST("/list", controllers.GetTheTvList)
	thetv.POST("/gallery/list", controllers.GetTheTvListByGalleryId)
	thetv.POST("/search", controllers.SearchTheTv)
	thetv.POST("/add", controllers.AddTheTv)
	thetv.POST("/sort", controllers.SortTheTv)

	// 剧集分集
	episode := r.Group("/v1/api/episode", auth.JWTAuth())
	episode.POST("/create", controllers.CreateEpisode)
	episode.POST("/update", controllers.UpdateEpisodeById)
	episode.POST("/delete", controllers.DeleteEpisodeById)
	episode.POST("/id", controllers.GetEpisodeById)
	episode.POST("/list", controllers.GetEpisodeList)
	episode.POST("/search", controllers.SearchEpisode)

	// 电视分季
	theseason := r.Group("/v1/api/theseason", auth.JWTAuth())
	theseason.POST("/create", controllers.CreateTheSeason)
	theseason.POST("/update", controllers.UpdateTheSeasonById)
	theseason.POST("/delete", controllers.DeleteTheSeasonById)
	theseason.POST("/id", controllers.GetTheSeasonById)
	theseason.POST("/list", controllers.GetTheSeasonList)
	theseason.POST("/search", controllers.SearchTheSeason)

	// 季信息
	season := r.Group("/v1/api/season", auth.JWTAuth())
	season.POST("/create", controllers.CreateSeason)
	season.POST("/update", controllers.UpdateSeasonById)
	season.POST("/delete", controllers.DeleteSeasonById)
	season.POST("/id", controllers.GetSeasonById)
	season.POST("/list", controllers.GetSeasonList)
	season.POST("/search", controllers.SearchSeason)

	lastepisodetoair := r.Group("/v1/api/lastepisodetoair", auth.JWTAuth())
	lastepisodetoair.POST("/create", controllers.CreateLastEpisodeToAir)
	lastepisodetoair.POST("/update", controllers.UpdateLastEpisodeToAirById)
	lastepisodetoair.POST("/delete", controllers.DeleteLastEpisodeToAirById)
	lastepisodetoair.POST("/id", controllers.GetLastEpisodeToAirById)
	lastepisodetoair.POST("/list", controllers.GetLastEpisodeToAirList)
	lastepisodetoair.POST("/search", controllers.SearchLastEpisodeToAir)

	nextepisodetoair := r.Group("/v1/api/nextepisodetoair", auth.JWTAuth())
	nextepisodetoair.POST("/create", controllers.CreateNextEpisodeToAir)
	nextepisodetoair.POST("/update", controllers.UpdateNextEpisodeToAirById)
	nextepisodetoair.POST("/delete", controllers.DeleteNextEpisodeToAirById)
	nextepisodetoair.POST("/id", controllers.GetNextEpisodeToAirById)
	nextepisodetoair.POST("/list", controllers.GetNextEpisodeToAirList)
	nextepisodetoair.POST("/search", controllers.SearchNextEpisodeToAir)

	networks := r.Group("/v1/api/networks", auth.JWTAuth())
	networks.POST("/create", controllers.CreateNetworks)
	networks.POST("/update", controllers.UpdateNetworksById)
	networks.POST("/delete", controllers.DeleteNetworksById)
	networks.POST("/id", controllers.GetNetworksById)
	networks.POST("/list", controllers.GetNetworksList)
	networks.POST("/search", controllers.SearchNetworks)

	// 影音库
	gallery := r.Group("/v1/api/gallery")
	gallery.POST("/create", auth.JWTAuthAdmin(), controllers.CreateGallery)
	gallery.POST("/update", auth.JWTAuthAdmin(), controllers.UpdateGalleryById)
	gallery.POST("/delete", auth.JWTAuthAdmin(), controllers.DeleteGalleryById)
	gallery.POST("/id", auth.JWTAuthAdmin(), controllers.GetGalleryById)
	gallery.POST("/list", auth.JWTAuth(), controllers.GetGalleryList)
	gallery.POST("/admin/list", auth.JWTAuthAdmin(), controllers.GetGalleryListAdmin)
	gallery.POST("/search", auth.JWTAuth(), controllers.SearchGallery)
	gallery.POST("/host", auth.JWTAuth(), controllers.GetGalleryHostByUid)

	// 刮削任务
	work := r.Group("/v1/api/work", auth.JWTAuthAdmin())
	work.POST("/create", controllers.CreateWork)
	work.POST("/renew", controllers.ReNewWork)
	work.POST("/update", controllers.UpdateWorkById)
	work.POST("/delete", controllers.DeleteWorkById)
	work.POST("/id", controllers.GetWorkById)
	work.POST("/list", controllers.GetWorkList)
	work.POST("/gallery/list", controllers.GetWorkListByGalleryId)
	work.POST("/search", controllers.SearchWork)

	// 刮削错误文件收集
	errfile := r.Group("/v1/api/errfile", auth.JWTAuthAdmin())
	errfile.POST("/create", controllers.CreateErrFile)
	errfile.POST("/update", controllers.UpdateErrFileById)
	errfile.POST("/delete", controllers.DeleteErrFileById)
	errfile.POST("/id", controllers.GetErrFileById)
	errfile.POST("/list", controllers.GetErrFileList)
	errfile.POST("/search", controllers.SearchErrFile)
	errfile.POST("/work/list", controllers.GetErrFilesByWorkId)
	errfile.POST("/ref/work/list", controllers.RefErrFilesByWorkId)
	errfile.POST("/ref/file/id", controllers.RefErrFileById)
	errfile.POST("/ref/themovie/id", controllers.RefErrTheMovieById)
	errfile.POST("/ref/thetv/id", controllers.RefErrTheTvById)
	errfile.POST("/ref/file/search", controllers.RefErrFileSearch)

	//收藏
	star := r.Group("/v1/api/star", auth.JWTAuth())
	star.POST("/create", controllers.CreateStar)
	star.POST("/update", controllers.UpdateStarById)
	star.POST("/delete", controllers.DeleteStarById)
	star.POST("/renew", controllers.ReNewStarByStar)
	star.POST("/id", controllers.GetStarById)
	star.POST("/list", controllers.GetStarList)
	star.POST("/search", controllers.SearchStar)
	star.POST("/data/list", controllers.GetStarDataList)

	// 点赞
	heart := r.Group("/v1/api/heart", auth.JWTAuth())
	heart.POST("/create", controllers.CreateHeart)
	heart.POST("/update", controllers.UpdateHeartById)
	heart.POST("/delete", controllers.DeleteHeartById)
	heart.POST("/renew", controllers.ReNewHeartByHeart)
	heart.POST("/id", controllers.GetHeartById)
	heart.POST("/list", controllers.GetHeartList)
	heart.POST("/data/list", controllers.GetHeartDataList)

	played := r.Group("/v1/api/played", auth.JWTAuth())
	played.POST("/create", controllers.CreatePlayed)
	played.POST("/update", controllers.UpdatePlayedById)
	played.POST("/delete", controllers.DeletePlayedById)
	played.POST("/renew", controllers.ReNewPlayedByPlayed)
	played.POST("/id", controllers.GetPlayedById)
	played.POST("/list", controllers.GetPlayedList)
	played.POST("/search", controllers.SearchPlayed)
	played.POST("/data/list", controllers.GetPlayedDataList)

	//客户端首屏api
	app := r.Group("/v1/api/app", auth.JWTAuth())
	app.POST("/index", controllers.AppIndex)

	// 阿里open
	aliOpen := r.Group("/v1/api/aliopen")
	aliOpen.POST("/video", controllers.AliOpenVideo)

	// 设置
	setting := r.Group("/v1/api/config", auth.JWTAuth())
	setting.POST("/save", auth.JWTAuthAdmin(), controllers.SaveConfig)
	setting.POST("/data", controllers.GetConfig)
	r.GET("/v1/api/configs", controllers.GetWebConfig)

	r.GET("/onelist/ping", func(c *gin.Context) {
		configData := config.GetConfig()
		configData.KeyDb = ""
		c.JSON(200, gin.H{"code": 200, "msg": "success", "data": configData})
	})

	r.GET("/t/p/*path", controllers.ImgServer)
	r.GET("/gallery/*path", controllers.GalleryImgServer)
	r.GET("/file/*path", controllers.FileServer)
	r.POST("/file/gallery/upload", controllers.FileUpload, auth.JWTAuthAdmin())
	r.Run(fmt.Sprintf(":%d", config.PORT))
}

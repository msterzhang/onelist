package thedb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/msterzhang/onelist/api/utils/dir"
)

var imgpath = "images"
var imgcdn = "https://tmdb-image-prod.b-cdn.net"

var dirs = []string{"w220_and_h330_face", "w440_and_h660_face", "w600_and_h900_bestv2", "w710_and_h400_multi_faces", "w227_and_h127_bestv2", "w1920_and_h1080_bestv2", "w355_and_h200_multi_faces"}

// 下载竖向海报图
var keys = []string{"w220_and_h330_face", "w440_and_h660_face", "w600_and_h900_bestv2"}

// 电视分季信息图片
var keysSeason = []string{"w220_and_h330_face"}

// 电视分集信息图片
var keysEpisode = []string{"w227_and_h127_bestv2", "w710_and_h400_multi_faces", "w1920_and_h1080_bestv2"}

// 下载横向海报图及背景图
var keysBackImge = []string{"w355_and_h200_multi_faces", "w1920_and_h1080_bestv2"}

// 初始化图片保存目录
func initDir() {
	for _, item := range dirs {
		imgsPath := imgpath + "/" + item
		if !dir.DirExists(imgsPath) {
			err := os.MkdirAll(imgsPath, os.ModePerm)
			if err != nil {
				log.Panic("创建图片保存文件夹失败!")
			}
		}
	}
}

// 下载电视剧及电影竖向海报图
func DownImages(id string) error {
	if len(id) == 0 {
		return nil
	}
	initDir()
	for _, key := range keys {
		url := fmt.Sprintf("%s/t/p/%s/%s", imgcdn, key, id)
		file := fmt.Sprintf("%s/%s/%s", imgpath, key, id)
		if dir.FileExists(file) {
			continue
		}
		err := Download(url, file)
		if err != nil {
			continue
		}
	}
	return nil
}

// 下载电视分季所需图片
func DownSeasonImages(id string) error {
	if len(id) == 0 {
		return nil
	}
	initDir()
	for _, key := range keysSeason {
		url := fmt.Sprintf("%s/t/p/%s/%s", imgcdn, key, id)
		file := fmt.Sprintf("%s/%s/%s", imgpath, key, id)
		if dir.FileExists(file) {
			continue
		}
		err := Download(url, file)
		if err != nil {
			continue
		}
	}
	return nil
}

// 下载电视分集所需图片
func DownEpisodeImages(id string) error {
	if len(id) == 0 {
		return nil
	}
	initDir()
	for _, key := range keysEpisode {
		url := fmt.Sprintf("%s/t/p/%s/%s", imgcdn, key, id)
		file := fmt.Sprintf("%s/%s/%s", imgpath, key, id)
		if dir.FileExists(file) {
			continue
		}
		err := Download(url, file)
		if err != nil {
			continue
		}
	}
	return nil
}

// 下载影人图片
func DownPersonImage(id string) error {
	if len(id) == 0 {
		return nil
	}
	initDir()
	url := fmt.Sprintf("%s/t/p/%s/%s", imgcdn, "w220_and_h330_face", id)
	file := fmt.Sprintf("%s/%s/%s", imgpath, "w220_and_h330_face", id)
	if dir.FileExists(file) {
		return nil
	}
	err := Download(url, file)
	if err != nil {
		return err
	}
	return nil
}

// 下载封面及大背景图
func DownBackImage(id string) error {
	if len(id) == 0 {
		return nil
	}
	initDir()
	for _, key := range keysBackImge {
		url := fmt.Sprintf("%s/t/p/%s/%s", imgcdn, key, id)
		file := fmt.Sprintf("%s/%s/%s", imgpath, key, id)
		if dir.FileExists(file) {
			continue
		}
		err := Download(url, file)
		if err != nil {
			continue
		}
	}
	return nil
}

// 下载图片
func Download(url string, fileName string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	client := http.Client{
		Timeout: timeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	io.Copy(file, resp.Body)
	return nil
}

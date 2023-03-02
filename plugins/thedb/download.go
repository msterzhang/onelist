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
var keys = []string{"w220_and_h330_face", "w355_and_h200_multi_faces", "w227_and_h127_bestv2"}

// 初始化图片保存目录
func initDir() {
	if !dir.DirExists(imgpath) {
		err := os.MkdirAll(imgpath, os.ModePerm)
		if err != nil {
			log.Panic("创建图片保存文件夹失败!")
		}
		for _, key := range keys {
			imgsPath := imgpath + "/" + key
			if !dir.DirExists(imgsPath) {
				err := os.MkdirAll(imgsPath, os.ModePerm)
				if err != nil {
					log.Panic("创建图片保存文件夹失败!")
				}
			}
		}
	}
}

// 下载各分辨率图片
func DownImages(id string) error {
	initDir()
	for _, key := range keys {
		url := fmt.Sprintf("https://image.tmdb.org/t/p/%s/%s", key, id)
		file := fmt.Sprintf("%s/%s/%s", imgpath, key, id)
		if dir.FileExists(file) {
			return nil
		}
		err := Download(url, file)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

// 下载影人图片
func DownPersonImage(id string) error {
	initDir()
	url := fmt.Sprintf("https://image.tmdb.org/t/p/%s/%s", "w220_and_h330_face", id)
	file := fmt.Sprintf("%s/%s/%s", imgpath, "w220_and_h330_face", id)
	if dir.FileExists(file) {
		return nil
	}
	err := Download(url, file)
	if err != nil {
		log.Println(err)
	}
	return nil
}

// 下载大背景图
func DownBackImage(id string) error {
	initDir()
	url := fmt.Sprintf("https://image.tmdb.org/t/p/%s/%s", "w1920_and_h1080_bestv2", id)
	file := fmt.Sprintf("%s/%s/%s", imgpath, "w1920_and_h1080_bestv2", id)
	if dir.FileExists(file) {
		return nil
	}
	err := Download(url, file)
	if err != nil {
		log.Println(err)
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

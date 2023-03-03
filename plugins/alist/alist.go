package alist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/config"
)

// 登录alist获取token
func AlistLogin(gallery models.Gallery) (string, error) {
	api := fmt.Sprintf("%s/api/auth/login", gallery.AlistHost)
	form := fmt.Sprintf(`{"username":"%s","password":"%s","otp_code":""}`, gallery.AlistUser, gallery.AlistPwd)
	req, err := http.NewRequest("POST", api, bytes.NewBufferString(form))
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", config.UA)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var data = AlistRspLogin{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	if data.Code == 200 {
		return data.Data.Token, nil
	}
	return "", errors.New(data.Message)
}

// 获取文件夹中文件及文件夹
func AlistFilesByPath(work models.Work, gallery models.Gallery, path string, Authorization string) ([]Content, error) {
	api := fmt.Sprintf("%s/api/fs/list", gallery.AlistHost)
	form := fmt.Sprintf(`{"path":"%s","password":"","page":1,"per_page":0,"refresh":%t}`, path, work.IsRef)
	req, err := http.NewRequest("POST", api, bytes.NewBufferString(form))
	if err != nil {
		return []Content{}, err
	}
	req.Header.Set("User-Agent", config.UA)
	req.Header.Set("Authorization", Authorization)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return []Content{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Content{}, err
	}
	var data = AListRspData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []Content{}, err
	}
	if data.Code == 200 {
		return data.Data.Content, nil
	}
	return []Content{}, errors.New(data.Message)
}

// 递归获取所有文件
func AlistList(work models.Work, gallery models.Gallery, path string, Authorization string, fileList []string) ([]string, error) {
	fs, err := AlistFilesByPath(work, gallery, path, Authorization)
	if err != nil {
		// 目录错误就重试一次
		fileList, err = AlistList(work, gallery, path, Authorization, fileList)
		if err != nil {
			if len(fileList) > 0 {
				return fileList, nil
			}
			return fileList, err
		}
	}
	for _, file := range fs {
		// 防止拼接path错误
		if path[len(path)-1:] != "/" {
			path += "/"
		}
		if file.IsDir {
			fileList, err = AlistList(work, gallery, path+file.Name+"/", Authorization, fileList)
			if err != nil {
				return fileList, err
			}
		} else {
			// 判断文件格式是否满足刮削条件
			fileAlistPath := "/d" + path + file.Name
			fileExt := filepath.Ext(fileAlistPath)
			if strings.Contains(config.VideoTypes, fileExt) {
				fileList = append(fileList, fileAlistPath)
			}
		}
	}
	return fileList, nil
}

// 根据目录获取alist中所有文件
func GetAlistFilesPath(work models.Work, gallery models.Gallery) ([]string, error) {
	fileList := []string{}
	path := work.Path
	Authorization, err := AlistLogin(gallery)
	if err != nil {
		return []string{}, err
	}
	// 防止提交host中含有"/"导致url拼接错误
	if gallery.AlistHost[len(gallery.AlistHost)-1:] == "/" {
		gallery.AlistHost = strings.TrimRight(gallery.AlistHost, "/")
	}
	return AlistList(work, gallery, path, Authorization, fileList)
}

// 刮削失败后修改文件名时候同时提交到alist修改
func AlistRnameFile(name string, errfile models.ErrFile) error {
	gallery := models.Gallery{}
	db := database.NewDb()
	err := db.Model(&models.Gallery{}).Where("gallery_uid = ?", errfile.GalleryUid).First(&gallery).Error
	if err != nil {
		return err
	}
	Authorization, err := AlistLogin(gallery)
	if err != nil {
		return err
	}
	api := fmt.Sprintf("%s/api/fs/rename", gallery.AlistHost)
	form := fmt.Sprintf(`{"path":"%s","name":"%s"}`, strings.ReplaceAll(errfile.File, "/d", ""), name)
	req, err := http.NewRequest("POST", api, bytes.NewBufferString(form))
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", config.UA)
	req.Header.Set("Authorization", Authorization)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if strings.Contains(string(body), "200") {
		return nil
	}
	return errors.New(string(body))
}

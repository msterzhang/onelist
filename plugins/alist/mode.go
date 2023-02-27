package alist

import "time"

type AlistRspLogin struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}

// alist文件体
type Content struct {
	Name        string    `json:"name"`
	ContentHash string    `json:"content_hash"`
	Size        int       `json:"size"`
	IsDir       bool      `json:"is_dir"`
	Modified    time.Time `json:"modified"`
	Sign        string    `json:"sign"`
	Thumb       string    `json:"thumb"`
	Type        int       `json:"type"`
}

// alist获取目录的
type AListRspData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Content  []Content `json:"content"`
		Total    int       `json:"total"`
		Readme   string    `json:"readme"`
		Write    bool      `json:"write"`
		Provider string    `json:"provider"`
	} `json:"data"`
}

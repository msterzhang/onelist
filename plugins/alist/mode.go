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

type AliOpenForm struct{
	File string `json:"file"`
	GalleryUid string `json:"gallery_uid"`
}


// alist 阿里云open资源
type AliOpenVideo struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data struct {
		DriveID string `json:"drive_id"`
		FileID string `json:"file_id"`
		VideoPreviewPlayInfo struct {
			Category string `json:"category"`
			LiveTranscodingSubtitleTaskList []struct {
				Language string `json:"language"`
				Status string `json:"status"`
				URL string `json:"url"`
			} `json:"live_transcoding_subtitle_task_list"`
			LiveTranscodingTaskList []struct {
				Stage string `json:"stage"`
				Status string `json:"status"`
				TemplateHeight int `json:"template_height"`
				TemplateID string `json:"template_id"`
				TemplateName string `json:"template_name"`
				TemplateWidth int `json:"template_width"`
				URL string `json:"url"`
			} `json:"live_transcoding_task_list"`
			Meta struct {
				Duration float64 `json:"duration"`
				Height int `json:"height"`
				Width int `json:"width"`
			} `json:"meta"`
		} `json:"video_preview_play_info"`
	} `json:"data"`
}

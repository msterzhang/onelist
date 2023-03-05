package extract

import (
	"errors"
	"path/filepath"
	"regexp"
	"strconv"
)

func removeEndingOne(s string) string {
	if len(s) > 0 && s[len(s)-1] == '1' {
		return s[:len(s)-1]
	}
	return s
}

// 过滤电影文件名
func ExtractMovieName(s string) string {
	// 删除发布年份和文件扩展名
	re := regexp.MustCompile(`\d{4}`)
	s = re.ReplaceAllString(s, "")

	// 删除括号及其内容
	re = regexp.MustCompile(`\s*\([^)]+\)`)
	s = re.ReplaceAllString(s, "")

	// 提取中文名称
	re = regexp.MustCompile(`[\p{Han}\d{1,2}]+`)
	matches := re.FindAllString(s, -1)
	if len(matches) > 0 {
		name := removeEndingOne(matches[0])
		return name
	}
	return s
}

// 根据文件名获取剧集季及集信息
func ExtractNumberWithFile(file string) (int, int, error) {
	p, err := filepath.Abs(file)
	if err != nil {
		return 0, 0, err
	}
	SeasonNumber := 0
	EpisodeNumber := 0
	fileName := filepath.Base(p)
	re := regexp.MustCompile(`[Ss](\d{1,2})[Ee](\d{1,2})`)
	match := re.FindStringSubmatch(fileName)
	if len(match) < 3 {
		return 0, 0, errors.New("get number error")
	}
	season := match[1]
	episode := match[2]
	SeasonNumber, err = strconv.Atoi(season)
	if err != nil {
		return 0, 0, err
	}
	EpisodeNumber, err = strconv.Atoi(episode)
	if err != nil {
		return 0, 0, err
	}
	return SeasonNumber, EpisodeNumber, nil
}

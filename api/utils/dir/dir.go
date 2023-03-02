package dir

import "os"

// 递归遍历目录中的文件
func GetFilesPath(path string, fileList []string) []string {
	fs, err := os.ReadDir(path)
	if err != nil {
		return fileList
	}
	for _, file := range fs {
		// 防止拼接path错误
		if path[len(path)-1:] != "/" {
			path += "/"
		}
		if file.IsDir() {
			fileList = GetFilesPath(path+file.Name()+"/", fileList)
		} else {
			fileList = append(fileList, path+file.Name())
		}
	}
	return fileList
}

// 获取目录中的所有文件
func GetFilesByPath(path string) []string {
	fileList := []string{}
	return GetFilesPath(path, fileList)
}

func DirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

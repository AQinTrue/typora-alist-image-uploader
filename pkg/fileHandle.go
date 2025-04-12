package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	Url   string `json:"url"`
	Path  string `json:"path"`
	Token string `json:"token"`
}

func SaveConfig(config Config, configPath string) bool {
	file, err := os.Create(configPath)
	if err != nil {
		return false
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // 格式化 JSON，设置缩进

	if err := encoder.Encode(config); err != nil {
		return false
	}

	return true
}

func ReadConfig(configPath string) (bool, Config) {
	var config Config

	file, err := os.Open(configPath)
	if err != nil {
		return false, config
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return false, config
	}

	return true, config
}

type Uploader struct {
	ImgPaths    []string
	NewImgNames []string
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (u *Uploader) Rename() {
	for _, path := range u.ImgPaths {
		imgExtension := filepath.Ext(path)
		imgExtension = strings.TrimPrefix(imgExtension, ".")

		timestampMillis := time.Now()
		millisSeconds := timestampMillis.Nanosecond() / 1e6
		timeFormat := timestampMillis.Format("2006_01_02_15_04_05")

		u.NewImgNames = append(u.NewImgNames, fmt.Sprintf("%s_%03d.%s", timeFormat, millisSeconds, imgExtension))
	}
}

func (u *Uploader) Uploader(url, path, token, imgPath, imgNewName string) string {

	fileContent, err := os.ReadFile(imgPath)
	if err != nil {
		fmt.Print("Upload Failed:\n", "读取图片错误.")
		return ""
	}

	body := bytes.NewBuffer(fileContent)

	headser := map[string]string{
		"Authorization":  token,
		"Content-Type":   "application/octet-stream",
		"Content-Length": fmt.Sprintf("%d", len(fileContent)),
		"file-path":      path + "/" + imgNewName,
	}

	request, err := http.NewRequest("PUT", url+"api/fs/put", body)

	if err != nil {
		fmt.Print("Upload Failed:\n", "创建请求失败.")
		return ""
	}

	for key, value := range headser {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		fmt.Print("Upload Failed:\n", "请求失败.")
		return ""
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return ""
		}

		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Print("Upload Failed:\n", "解析失败.")
			return ""
		}

		serverUrl := url + "d/" + path + "/" + imgNewName + "\n"
		return serverUrl
	} else {
		fmt.Print("Upload Failed:\n", "请求错误.")
		return ""
	}
}

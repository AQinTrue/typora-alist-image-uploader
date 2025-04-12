package pkg

import (
	"flag"
)

type SourceData struct {
	Type     string
	Url      string
	Path     string
	Token    string
	ImgPaths []string
}

type DataCode struct {
	Upload bool
	Update bool
	Err    bool
}

func getParser() SourceData {
	typ := flag.String("type", "", "模式类型 [uploadImg, updateConfig]")
	url := flag.String("url", "", "alist的网址, 例如: https://example.com")
	path := flag.String("path", "", "保存的路径, 例如: dir1/dir2")
	token := flag.String("token", "", "管理员的token")

	flag.Parse()

	imgPaths := flag.Args()

	return SourceData{
		Type:     *typ,
		Url:      *url,
		Path:     *path,
		Token:    *token,
		ImgPaths: imgPaths,
	}
}

func CheckCode() (DataCode, SourceData) {
	parse := getParser()
	switch parse.Type {
	case "uploadImg":
		if len(parse.ImgPaths) == 0 {
			return DataCode{Err: true}, parse
		}
		return DataCode{Upload: true}, parse

	case "updateConfig":
		if parse.Url == "" || parse.Path == "" || parse.Token == "" {
			return DataCode{Err: true}, parse
		}
		return DataCode{Update: true}, parse

	default:
		return DataCode{Err: true}, parse
	}
}

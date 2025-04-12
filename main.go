package main

import (
	"fmt"
	"os"
	"path/filepath"
	"taiu/pkg"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Upload Failed:", err)
		return
	}

	dir := filepath.Dir(exePath)
	configPath := filepath.Join(dir, "Taiu.json")

	code, parse := pkg.CheckCode()
	if code.Upload {
		code, config := pkg.ReadConfig(configPath)
		if code {
			pkg.NewDataCrypto().Decrypt(config.Url)
			url := pkg.NewDataCrypto().Decrypt(config.Url)
			upath := pkg.NewDataCrypto().Decrypt(config.Path)
			token := pkg.NewDataCrypto().Decrypt(config.Token)

			uploader := &pkg.Uploader{
				ImgPaths: parse.ImgPaths,
			}
			uploader.Rename()

			result := "Upload Success:\n"

			for i, imgPath := range uploader.ImgPaths {
				result += uploader.Uploader(url, upath, token, imgPath, uploader.NewImgNames[i])
			}

			fmt.Print(result)

		} else {
			fmt.Println("Upload Failed:", "读取配置失败.\n"+configPath)
		}
	}

	if code.Update {
		url := pkg.NewDataCrypto().Encrypt(parse.Url)
		path := pkg.NewDataCrypto().Encrypt(parse.Path)
		token := pkg.NewDataCrypto().Encrypt(parse.Token)
		config := pkg.Config{
			Url:   url,
			Path:  path,
			Token: token,
		}
		if pkg.SaveConfig(config, configPath) {
			fmt.Println("保存成功.")
		} else {
			fmt.Println("保存失败.")
		}
	}

	if code.Err {
		fmt.Println("taiu -h 获取帮助.")
		os.Exit(1)
	}
}

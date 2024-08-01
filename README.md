# Typora-Alist-Image-Uploader
## 这是一个简单的工具，用于方便地将Typora编辑器中的图片上传至Alist云存储服务。通过这个工具，用户可以轻松管理Typora文档中的图片资源，无需手动上传和管理文件。

## 特点:
- 一键上传: 直接从Typora编辑器上传图片至Alist。
- 自动替换: 自动替换文档中的图片链接为已上传的图片URL。
- 高效工作流: 提高写作和编辑效率，简化图片管理流程。
## 使用方法:
- 下载解压。
- 查看参数
```shell
taiu-win64-v1.0.0.exe -h
```
- 初始化配置
```shell
taiu-win64-v1.0.0.exe --type uconfig --url https://# --path taiu/img --token alist-token
```
- Typora配置
```shell
D:\PATH\TAIU.exe --type upload --upload_args 
```
- 在Typora中插入图片。
- 查看上传结果。

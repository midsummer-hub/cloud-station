### 文件中转站

##### 1、配置个人的云服务秘钥的默认参数
路径：cmd\client\cmd\upload.go
```go
defaultBuckName = "xx-x1"
defaultEndpoint = "http://oss-cn-xxxx.aliyuncs.com"
defaultALIAK    = "xxxxxxxx"
defaultALISK    = "xxxxxxxx"
```
##### 2、跨平台编译
Mac 下编译 Linux 和 Windows平台 64位 可执行程序：
```go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/client/main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build cmd/client/main.go
```
Linux 下编译 Mac 和 Windows 平台64位可执行程序：
```go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build cmd/client/main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build cmd/client/main.go
```
Windows下编译Mac平台64位可执行程序：
```go
CGO_ENABLED=0 SET GOOS=darwin GOARCH=amd64 go build cmd/client/main.go
CGO_ENABLED=0 SET GOOS=linux GOARCH=amd64 go build cmd/client/main.go
```
##### 3、简单使用
帮助信息 -h
```go
./main.exe upload -h

>>> output
上传文件到中转站

Usage:
  cloud-station-cli upload [flags]

Flags:
  -e, --bucket_endpoint string   upload oss endpoint (default "http://oss-cn-xxxx.aliyuncs.com")
  -b, --bucket_name string       upload oss bucket name (default "xx-x1")
  -f, --file_path string         upload file path
  -h, --help                     help for upload

Global Flags:
  -i, --ali_access_id string    the ali oss access id (default "xxxxxxxx")
  -k, --ali_access_key string   the ali oss access key (default "xxxxxxxx")
  -p, --oss_provider string     the oss provider aliyun (default "aliyun")
  -v, --version                 the cloud-station-cli version
```
指定文件上传
```go
./main.exe upload -f text.txt
>>> output
请输入阿里云SK:  ******************************
开始上传: 100% [==============================] (5.271 kB/s)
下载链接：http://cloud-x1.oss-cn-beijing.aliyuncs.com/test.txt?Expires=1635533897&OSSAccessKeyId=xx&Signature=xx
注意：文件下载有效期1天,中转站保存时间3天,请及时下载.
```

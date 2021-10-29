package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/china-sd/cloud-station/store"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func NewUploader(endpoint, accessKeyID, accessKeySecret string) (store.Uploader, error) {
	uploader := &aliyun{
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		listener:        NewListener(),
	}
	if err := uploader.Validate(); err != nil {
		return nil, err
	}

	return uploader, nil
}

type aliyun struct {
	Endpoint        string `validate:"required,url"`
	AccessKeyID     string `validate:"required"`
	AccessKeySecret string `validate:"required"`
	listener        oss.ProgressListener
}

func (a *aliyun) Validate() error {
	return validate.Struct(a)
}

func (a *aliyun) UploadFile(bucketName, objectKey, localFilePath string) error {
	if objectKey == "" || localFilePath == "" {
		return fmt.Errorf("upload file missed.")
	}
	client, err := oss.New(a.Endpoint, a.AccessKeyID, a.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(localFilePath, localFilePath, oss.Progress(a.listener))
	if err != nil {
		return err
	}

	singleURL, err := bucket.SignURL(localFilePath, oss.HTTPGet, 6*60*24)
	if err != nil {
		return fmt.Errorf("sign file download url error , %s.\n", err)
	}
	fmt.Printf("\n下载链接：%s\n", singleURL)
	fmt.Printf("注意：文件下载有效期1天,中转站保存时间3天,请及时下载.")

	return nil
}

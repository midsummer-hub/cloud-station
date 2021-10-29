package aliyun_test

import (
	"github.com/china-sd/cloud-station/store/provider/aliyun"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	endpoint        = "http://oss-cn-beijing.aliyuncs.com"
	accessKeyID     = "xxxxxxxx"
	accessKeySecret = "xxxxxxxx"
	bucketName      = "xx-x1"
	filePath        = ""
)

func TestUploadFile(t *testing.T) {
	should := assert.New(t)

	uploader, err := aliyun.NewUploader(endpoint, accessKeyID, accessKeySecret)
	if should.NoError(err) {
		err = uploader.UploadFile(bucketName, filePath, filePath)
		should.NoError(err)
	}
}

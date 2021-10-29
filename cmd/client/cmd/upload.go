package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/china-sd/cloud-station/store"
	"github.com/china-sd/cloud-station/store/provider/aliyun"
	"github.com/spf13/cobra"
	"path"
	"time"
)

var (
	defaultBuckName = "xx-x1"
	defaultEndpoint = "http://oss-cn-xxxx.aliyuncs.com"
	defaultALIAK    = "xxxxxxxx"
	defaultALISK    = "xxxxxxxx"
)

var (
	buckName       string
	uploadFilePath string
	bucketEndpoint string
)

// uploadCmd represents the start command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传文件到中转站",
	Long:  `上传文件到中转站`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := getUploader()
		if err != nil {
			return err
		}

		if uploadFilePath == "" {
			return fmt.Errorf("upload file path is missing")
		}
		day := time.Now().Format("20060102")
		fn := path.Base(uploadFilePath)
		ok := fmt.Sprintf("%s/%s", day, fn)
		err = p.UploadFile(buckName, ok, uploadFilePath)
		if err != nil {
			return err
		}
		return nil
	},
}

func getUploader() (store.Uploader, error) {
	switch ossProvider {
	case "aliyun":
		prompt := &survey.Password{
			Message: "请输入阿里云SK: ",
		}
		survey.AskOne(prompt, &aliAccessKey)
		return aliyun.NewUploader(bucketEndpoint, aliAccessID, aliAccessKey)
	case "qcloud":
		return nil, fmt.Errorf("not impl")
	case "minio":
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("unknown uploader %s", ossProvider)
	}
}

func init() {
	uploadCmd.PersistentFlags().StringVarP(&uploadFilePath, "file_path", "f", "", "upload file path")
	uploadCmd.PersistentFlags().StringVarP(&buckName, "bucket_name", "b", defaultBuckName, "upload oss bucket name")
	uploadCmd.PersistentFlags().StringVarP(&bucketEndpoint, "bucket_endpoint", "e", defaultEndpoint, "upload oss endpoint")
	RootCmd.AddCommand(uploadCmd)
}

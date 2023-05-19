package processors

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pokyux/Assistant/conf"
)

var bucket *oss.Bucket

func InitOSS() {
	client, err := oss.New(conf.OSSEndpoint, conf.OSSAccessKeyId, conf.OSSAccessKeySecret)
	if err != nil {
		panic(err)
	}

	bucket, err = client.Bucket(conf.OSSBucketName)
	if err != nil {
		panic(err)
	}
}

func UploadToOSS(rcvd *tgbotapi.Message, rply *tgbotapi.MessageConfig) {
	rply.Text = "Uploaded to oss."

}

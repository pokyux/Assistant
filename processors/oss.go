package processors

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pokyux/Assistant/conf"
	"github.com/pokyux/Assistant/global"
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
	if !IsNormalUser(rcvd.From.ID) {
		rply.Text = "Permission denied."
		return
	}

	if rcvd.Document == nil {
		rply.Text = "No document. Failed."
		return
	}

	file, err := global.Bot.GetFile(tgbotapi.FileConfig{FileID: rcvd.Document.FileID})
	if err != nil {
		rply.Text = "No document. Failed."
		return
	}

	link := file.Link(conf.TGBotToken)
	resp, err := http.Get(link)
	if err != nil {
		rply.Text = "Cannot get file from TG server."
		return
	}

	linkSplt := strings.Split(link, ".")
	fileExt := linkSplt[len(linkSplt)-1]
	filePath := fmt.Sprintf("botupld/%d/%s.%s", rcvd.From.ID, file.FileID, fileExt)
	err = bucket.PutObject(filePath, resp.Body)
	if err != nil {
		rply.Text = "Uploading failed."
		return
	}

	rply.Text = conf.OSSDomain + filePath

	// for index, photo := range rcvd.Photo {
	// 	file, err := global.Bot.GetFile(tgbotapi.FileConfig{FileID: photo.FileID})
	// 	if err != nil {
	// 		rply.Text += fmt.Sprintf("\n%d. Uploading failed.", index)
	// 		continue
	// 	}

	// 	link := file.Link(conf.TGBotToken)
	// 	resp, err := http.Get(link)
	// 	if err != nil {
	// 		rply.Text += fmt.Sprintf("\n%d. Uploading failed.", index)
	// 		continue
	// 	}

	// 	linkSplt := strings.Split(link, ".")
	// 	fileExt := linkSplt[len(linkSplt)-1]
	// 	filePath := fmt.Sprintf("/botupld/%d/%s.%s", rcvd.From.ID, file.FileID, fileExt)
	// 	err = bucket.PutObject(filePath, resp.Body)
	// 	if err != nil {
	// 		rply.Text += fmt.Sprintf("\n%d. Uploading failed.", index)
	// 		continue
	// 	}

	// 	resultLink := conf.OSSEndpoint + filePath
	// 	rply.Text += fmt.Sprintf("\n%d. %s", index, resultLink)
	// }
}

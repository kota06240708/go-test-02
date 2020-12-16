package util

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// s3に画像をアップロード
func UploadToS3(imageBase64 string) (url chan string, err chan error) {
	url = make(chan string)
	err = make(chan error)

	go func(imageBase64 string) {

		// 環境変数からS3Credential周りの設定を取得
		AWS_S3_BUCKET := os.Getenv("AWS_S3_BUCKET")
		AWS_ACCESS_KEY := os.Getenv("AWS_ACCESS_KEY")
		AWS_SECRET_KEY := os.Getenv("AWS_SECRET_KEY")
		AWS_S3_REGION := os.Getenv("AWS_S3_REGION")

		// awsの接続情報を設定
		sess := session.Must(session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY, AWS_SECRET_KEY, ""),
			Region:      aws.String(AWS_S3_REGION),
		}))

		uploader := s3manager.NewUploader(sess)

		// base64のデータを削除
		rep := regexp.MustCompile(`^data:\w+\/\w+;base64,`)
		str := rep.ReplaceAllString(imageBase64, "")

		// 拡張子を取得
		fileExtensionRep := regexp.MustCompile(`jpg|jpeg|gif|png`)
		fileExtension := fileExtensionRep.FindString(imageBase64)

		if fileExtension == "" {
			err <- fmt.Errorf("Error %s", "non Image")
			return
		}

		// // ContentTypeを取得 (今回は使わないのでコメントアウト)
		// removeData := strings.TrimLeft(imageBase64, `data:`)
		// contentTypeRep := regexp.MustCompile(`^image/\w+;`)
		// contentType := contentTypeRep.FindString(removeData)
		// fmt.Println(strings.TrimRight(contentType, ";"))

		// base64デコード
		data, decodeErr := base64.StdEncoding.DecodeString(str)

		if decodeErr != nil {
			err <- fmt.Errorf("Error decodeErr %s", decodeErr)
			return
		}

		// デコードしたbase64をバッファーに格納
		wb := new(bytes.Buffer)
		_, errImage := wb.Write(data)

		// base64が問題ないかチェック
		if errImage != nil {
			err <- fmt.Errorf("Error %s", errImage)
			return
		}

		// ファイル名を作成
		timeStanp := strconv.FormatInt(time.Now().Unix(), 10)
		fileName := timeStanp + "." + fileExtension

		// awsにアップロード
		res, errUpdate := uploader.Upload(&s3manager.UploadInput{
			Bucket:      aws.String(AWS_S3_BUCKET),
			Key:         aws.String(fileName),
			Body:        wb,
			ContentType: aws.String("image/" + fileExtension),
		})

		// エラーハンドリング
		if errUpdate != nil {
			if errAwse, ok := errUpdate.(awserr.Error); ok && errAwse.Code() == request.CanceledErrorCode {
				err <- fmt.Errorf("Error %s", "Upload TimuOut")
				return
			} else {
				err <- fmt.Errorf("Error %s", "Upload Failed")
				return
			}
		}

		// チャネルに登録
		url <- res.Location
	}(imageBase64)

	return url, err
}

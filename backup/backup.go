package backup

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/soxft/mysql-backuper/config"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func ToCos(flocation, filename string) error {

	bucketURL, _ := url.Parse("https://" + config.C.Cos.Bucket + ".cos." + config.C.Cos.Region + ".myqcloud.com")
	b := &cos.BaseURL{BucketURL: bucketURL}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.C.Cos.Secret.Id,
			SecretKey: config.C.Cos.Secret.Key,
		},
	})

	// check if path ends with "/"
	if config.C.Cos.Path[len(config.C.Cos.Path)-1:] != "/" {
		config.C.Cos.Path += "/"
	}
	// not start with "/"
	if config.C.Cos.Path[0:1] == "/" {
		config.C.Cos.Path = config.C.Cos.Path[1:]
	}

	ctx := context.Background()
	key := config.C.Cos.Path + filename
	log.Println(key)

	// 预签名
	opt := &cos.PresignedURLOptions{
		Query:  &url.Values{},
		Header: &http.Header{},
	}

	presignedURL, err := client.Object.GetPresignedURL(ctx, http.MethodPut, key, config.C.Cos.Secret.Id, config.C.Cos.Secret.Key, time.Minute*10, opt)
	if err != nil {
		return err
	}
	log.Println(presignedURL)
	// read file from local
	f, err := os.Open(flocation)
	if err != nil {
		return err
	}
	defer f.Close()

	req, err := http.NewRequest(http.MethodPut, presignedURL.String(), f)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	// out response
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(data))
	return nil
}

package backup

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/soxft/db-backuper/config"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func ToCos(flocation, filename string) (string, error) {

	bucketURL, _ := url.Parse("https://" + config.Cos.Bucket + ".cos." + config.Cos.Region + ".myqcloud.com")
	b := &cos.BaseURL{BucketURL: bucketURL}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.Cos.Secret.Id,
			SecretKey: config.Cos.Secret.Key,
		},
	})

	// check Bucket exist
	ok, err := client.Bucket.IsExist(context.Background())

	if err != nil {
		return "", err
	} else if !ok {
		return "", errors.New("bucket does not exist")
	}

	// check if path ends with "/"
	if config.Cos.Path[len(config.Cos.Path)-1:] != "/" {
		config.Cos.Path += "/"
	}
	// not start with "/"
	if config.Cos.Path[0:1] == "/" {
		config.Cos.Path = config.Cos.Path[1:]
	}

	// upload
	remoteFullPath := config.Cos.Path + filename
	_, err = client.Object.PutFromFile(context.Background(), remoteFullPath, flocation, nil)
	if err != nil {
		return "", err
	}
	return remoteFullPath, nil
}

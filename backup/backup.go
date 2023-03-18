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

	bucketURL, _ := url.Parse("https://" + config.C.Cos.Bucket + ".cos." + config.C.Cos.Region + ".myqcloud.com")
	b := &cos.BaseURL{BucketURL: bucketURL}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.C.Cos.Secret.Id,
			SecretKey: config.C.Cos.Secret.Key,
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
	if config.C.Cos.Path[len(config.C.Cos.Path)-1:] != "/" {
		config.C.Cos.Path += "/"
	}
	// not start with "/"
	if config.C.Cos.Path[0:1] == "/" {
		config.C.Cos.Path = config.C.Cos.Path[1:]
	}

	// upload
	remoteFullPath := config.C.Cos.Path + filename
	_, err = client.Object.PutFromFile(context.Background(), remoteFullPath, flocation, nil)
	if err != nil {
		return "", err
	}
	return remoteFullPath, nil
}

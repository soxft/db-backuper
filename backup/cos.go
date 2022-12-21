package backup

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/soxft/db-backuper/config"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// ToCos
// @Param flocation: local file location
// @Param dbname: database name
func ToCos(flocation, dbname string) (string, error) {
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

	// upload
	remotePath := config.Cos.Path + dbname + "/"
	remoteFullPath := remotePath + flocation[len(config.Local.Dir):]
	_, err = client.Object.PutFromFile(context.Background(), remoteFullPath, flocation, nil)
	if err != nil {
		return "", err
	}

	_ = removeCosMax(remotePath, config.Cos.MaxFileNum)
	return remoteFullPath, nil
}

// removeMax 删除 cos 上 指定路径下超过 max 个文件
func removeCosMax(remotePath string, max int) error {
	bucketURL, _ := url.Parse("https://" + config.Cos.Bucket + ".cos." + config.Cos.Region + ".myqcloud.com")
	b := &cos.BaseURL{BucketURL: bucketURL}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.Cos.Secret.Id,
			SecretKey: config.Cos.Secret.Key,
		},
	})

	opt := &cos.BucketGetOptions{
		Prefix:  "backups/timeletters",
		MaxKeys: 1000,
	}
	res, _, err := client.Bucket.Get(context.Background(), opt)
	if err != nil {
		return err
	}

	num := len(res.Contents)
	for _, v := range res.Contents {
		if num > max {
			_, _ = client.Object.Delete(context.Background(), v.Key)
			log.Println("remove cos file:", v.Key)
			num--
		} else {
			break
		}
	}
	return nil
}

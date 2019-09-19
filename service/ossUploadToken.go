package service

import (
	"btube/serializer"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

//OSSUploadTokenService use to get OSS upload singnURL and get video singnURL.
type OSSUploadTokenService struct {
	FileName string `json:"filename" form:"filename"`
}

//GetSignURLForVideo get the url with sign from OSS.
func (tokenSrv *OSSUploadTokenService) GetSignURLForVideo() *serializer.Response {
	endpoint := os.Getenv("OSS_ENDPOINT")
	accessKey := os.Getenv("OSS_ACCESS_KEY")
	bucketName := os.Getenv("OSS_BUCKET_NAME")
	ossScrecct := os.Getenv("OSS_SECRECT")
	client, err := oss.New(endpoint, accessKey, ossScrecct, oss.Timeout(10, 10000))
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "can't get token for upload.",
		}
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "can't get token for upload.",
		}
	}

	ext := filepath.Ext(tokenSrv.FileName)
	OssID := "/b-tube/video" + uuid.Must(uuid.NewRandom()).String() + ext
	signGetURL, err := bucket.SignURL(OssID, oss.HTTPGet, 10000, nil)
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "error when try to get oss get link.",
		}
	}
	signPutURL, err := bucket.SignURL(OssID, oss.HTTPPut, 315360000, nil)
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "error when try to get oss get link.",
		}
	}

	return &serializer.Response{
		StatusCode: 0,
		Data: map[string]string{
			"OssID": OssID,
			"put":   signPutURL,
			"get":   signGetURL,
		},
	}
}

//GetSignURLForAvatar get the url with sign from OSS.
func (tokenSrv *OSSUploadTokenService) GetSignURLForAvatar() *serializer.Response {
	endpoint := os.Getenv("OSS_ENDPOINT")
	accessKey := os.Getenv("OSS_ACCESS_KEY")
	bucketName := os.Getenv("OSS_BUCKET_NAME")
	ossScrecct := os.Getenv("OSS_SECRECT")
	client, err := oss.New(endpoint, accessKey, ossScrecct, oss.Timeout(10, 92400))
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "can't get token for upload.",
		}
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "can't get token for upload.",
		}
	}

	ext := filepath.Ext(tokenSrv.FileName)
	OssID := "/b-tube/avatar" + uuid.Must(uuid.NewRandom()).String() + ext
	signGetURL, err := bucket.SignURL(OssID, oss.HTTPGet, 10000, nil)
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "error when try to get oss get link.",
		}
	}
	signPutURL, err := bucket.SignURL(OssID, oss.HTTPPut, 10000, nil)
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "error when try to get oss get link.",
		}
	}

	return &serializer.Response{
		StatusCode: 0,
		Data: map[string]string{
			"OssID": OssID,
			"put":   signPutURL,
			"get":   signGetURL,
		},
	}
}

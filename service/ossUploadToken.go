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
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "can't get token for upload for oss configure error.",
		}
	}

	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "can't get token for upload for oss configure error.",
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

	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "can't get token for upload for oss configure error.",
		}
	}

	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "can't get token for upload for oss configure error.",
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

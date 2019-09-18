package service

//OSSUploadTokenService use to get OSS upload singnURL and get video singnURL.
type OSSUploadTokenService struct {
	FileName string `json:"filename" form:"filename"`
}

//GetSignURL get the url with sign from OSS.
func (tokenSrv *OSSUploadTokenService) GetSignURL() {

}

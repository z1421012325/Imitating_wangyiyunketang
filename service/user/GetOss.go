package user

import (
	"demos/serialize"
	"demos/util"
	"path"
	"strings"
)

type GetOssService struct {
	// 多文件上传,文件名以分号隔离
	FileNames string   `json:"filenames" form:"filenames" binding:"required"`
	// 文件类型,根据类型获取token
	Type      string   `json:"type"      form:"type"      binding:"required"`
}



func (service *GetOssService)GetOss()*serialize.Response  {

	var tokens interface{}

	if service.Type == "image" {
		tokens = util.GetImageToken(service.FileNames)
	}else if service.Type == "video" {
		tokens = util.GetVideoToken(service.FileNames)
	}else if service.Type == "other" {
		tokens = util.GetImageToOtherToken(service.FileNames)
	}

	return serialize.Res(tokens,"")
}

















func (service *GetOssService) VerifyVideoSuffix()bool {
	suffixs := []string{".mp4",".avi"}

	files := strings.Split(service.FileNames,",")
	for _,file := range files{
		filesuffix := path.Ext(file)
		for _ , suffix := range suffixs{
			if filesuffix != suffix{
				return false
			}
		}
	}
	return true
}
func (service *GetOssService) VerifyImageSuffix()bool {
	suffixs := []string{".jpg",".png"}

	files := strings.Split(service.FileNames,",")
	for _,file := range files{
		filesuffix := path.Ext(file)
		for _ , suffix := range suffixs{
			if filesuffix != suffix{
				return false
			}
		}
	}
	return true
}

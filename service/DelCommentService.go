package service

import (
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type DelCommentService struct {

}


func (service *DelCommentService)DelComment(c *gin.Context)*serialize.Response{
	var data model.CurriculumComment


	return serialize.Res(nil,"")
}
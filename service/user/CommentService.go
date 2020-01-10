package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)


type CurriculumCommentData struct {
	Result 	[]model.CurriculumComment		`json:"comments"`
	Total 	int								`gorm:"column:total" json:"total"`
}

func CommentService(c *gin.Context) *serialize.Response{

	cid := c.Param("cid")
	start,num := service.PagingQuery(c)

	var data CurriculumCommentData
	sql := "select " +
				"u.u_id,u.nickename,u.r_id," +
				"cm.c_id,cm.number,cm.comment,cm.create_at " +
//				"(select count(*) from curriculum_comments as cm where cm.c_id = ?) as total " +
			"from " +
				"curriculum_comments as cm join users as u " +
			"on " +
				"u.u_id = cm.u_id " +
			"where " +
				"cm.c_id = ? limit ?,?"

	DB.DB.Raw(sql,cid,start,num).Scan(&data.Result)
	DB.DB.Model(&model.CurriculumComment{}).Where("c_id = ?",cid).Count(&data.Total)

	return serialize.Res(data,"")
}



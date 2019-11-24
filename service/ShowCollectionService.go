package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"fmt"
	"github.com/gin-gonic/gin"
)

type usercoll struct {
	Result []model.Curriculums				`json:"result"`
	Total int								`gorm:"column:total" json:"total"`
}

func ShowCollectionService(c *gin.Context)*serialize.Response{
	uid := GetUserId(c)
	start,size := pagingQuery(c)

	var data usercoll
	sql := "select " +
				"cc.c_id,cc.u_id,cc.c_name,cc.price,cc.c_image,cc.create_at " +
			"from " +
				"use_collections as uc join curriculums as cc " +
			"on " +
				"cc.c_id = uc.c_id " +
			"where " +
				"uc.u_id = ? and uc.delete_at is null " +
			"order by uc.create_at " +
				"limit ?,?"

	DB.DB.Raw(sql,uid,start,size).Scan(&data.Result)
	DB.DB.Model(&model.UseCollections{}).Where("u_id = ?",uid).Count(&data.Total)

	fmt.Println(uid,start,size)
	fmt.Println("\n")
	fmt.Println(data)

	return serialize.Res(data,"")
}

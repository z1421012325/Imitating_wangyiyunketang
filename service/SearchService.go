package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"

	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SearchServiceData1 struct {
	model.Curriculums
	Number string 					`gorm:"column:number" json:"number"`	// todo 坑 avg映射的字段类型是string而不是int
	Count  string					`gorm:"column:count" json:"count"`
}
type SearchServiceData struct {
	Result []SearchServiceData1 `json:"datas"`
	Total  int                  `gorm:"column:total" json:"total"`
}


func SearchService(c *gin.Context) *serialize.Response{
	// 最新,最热门,all
	// 标签,课程标题模糊搜索
	start,end := pagingQuery(c)

	key := c.Query("key")		// 查询字段
	order := c.Query("order")	// 排序字段
	searchSql := ""
	// totalSql := ""

	if order == "new" {					// 最新 字段创建时间查询
		//"select c.c_id,c.u_id,c.c_name,c.price,c.c_image,c.create_at,avg(cc.number) as number from curriculum_comments as cc join curriculums as c on c.c_id = cc.c_id join crriculu_tag as ct on c.c_id = ct.c_id join tags as t on ct.t_id = t.t_id where c.c_name like '%re%' or t.t_name like '%re%' and c.create_at is null group by c.c_id order by c.create_at asc limit 0,5"
		searchSql = "select " +
					"c.c_id,c.u_id,c.c_name,c.price,c.c_image,c.create_at,avg(cc.number)as number " +
				"from " +
					"curriculum_comments as cc " +
				"join " +
					"curriculums as c " +
				"on " +
					"c.c_id = cc.c_id " +
				"join " +
					"crriculu_tag as ct " +
				"on " +
					"c.c_id = ct.c_id join tags as t " +
				"on " +
					"ct.t_id = t.t_id " +
				"where " +
					"c.c_name like '%"+key+"%' or t.t_name like '%"+key+"%' " +	 // 不能使用gorm的字段拼接,只能手动 去除 繁殖sql注入
					"and c.create_at is null " +
					"group by c.c_id order by c.create_at asc limit ?,?"

	}else if order == "hot" {			// 最热门 学习人数最多的

		//select c.c_id,c.u_id,c.c_name,c.price,c.c_image,c.create_at,avg(cc.number)as number,count(sp.u_id)as count from curriculum_comments as cc join curriculums as c on cc.c_id = c.c_id join shopping_carts as sp on c.c_id = sp.c_id join crriculu_tag as ct on c.c_id = ct.c_id join tags as t on ct.t_id = t.t_id where c.c_name like '%re%' or t.t_name like '%r%' and c.delete_at is null group by c.c_id,sp.u_id order by count desc
		searchSql = "select " +
					"c.c_id,c.u_id,c.c_name,c.price,c.c_image,c.create_at,avg(cc.number)as number,count(sp.u_id)as count " +
				"from " +
					"curriculum_comments as cc join curriculums as c " +
				"on " +
					"cc.c_id = c.c_id join shopping_carts as sp " +
				"on " +
					"c.c_id = sp.c_id join crriculu_tag as ct " +
				"on " +
					"c.c_id = ct.c_id join tags as t " +
				"on " +
					"ct.t_id = t.t_id " +
				"where " +
					"c.c_name like '%"+key+"%' or t.t_name like '%"+key+"%' " +
				"and c.delete_at is null group by c.c_id,sp.u_id order by count desc " +
					"limit ?,?"

	}else {								// 全部

		//  select c.c_id,c.u_id,c.c_name,c.price,c.c_image,c.create_at,avg(cc.number)as number from curriculum_comments as cc join curriculums as c on cc.c_id = c.c_id join crriculu_tag as ct on c.c_id = ct.c_id join tags as t on ct.t_id = t.t_id where c.c_name like '%re%' or t.t_name like '%r%' and c.delete_at is null group by c.c_id
		searchSql = "select " +
					"c.c_id,c.u_id,c.c_name,c.price,c.c_image,c.create_at,avg(cc.number)as number " +
				"from " +
					"curriculum_comments as cc join curriculums as c on cc.c_id = c.c_id " +
				"join " +
					"crriculu_tag as ct on c.c_id = ct.c_id " +
				"join " +
					"tags as t on ct.t_id = t.t_id " +
				"where " +
					"c.c_name like '%"+key+"%' or t.t_name like '%"+key+"%' and c.delete_at is null " +
				"group by c.c_id limit ?,?"

	}

	fmt.Printf("\n--------------\n sql语句为%s \n search查询词为%s \n ---------\n",searchSql,key)

	var data SearchServiceData
	DB.DB.Raw(searchSql,start,end).Scan(&data.Result)

	// 序列化数据
	value ,err := json.Marshal(data.Result)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(string(value))

	// 查询数据的总和
	// DB.DB.Raw(totalSql)

	return serialize.Res(data,"")
}

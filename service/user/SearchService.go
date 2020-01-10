package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)

type SearchServiceData1 struct {
	model.Curriculums
	Number string 					`gorm:"column:number" json:"number"`  	// 坑 集合查询映射的字段类型是string而不是int
	Count  string					`gorm:"column:count" json:"count"`
}
type SearchServiceData struct {
	Result []SearchServiceData1 	`json:"result"`
	Total  string                  	`gorm:"column:total" json:"total"`
}

// todo 设置redis 缓存
// todo 表的设计有缺陷,查询语句出问题了如果添加join 评论表 那么课程没有评价则不会查询出来,有则会查询出,
//  解决办法 not join 评论表,但是评价平均分得不到,
//  第二个则是在建课程时在评论表插入评论和必须添加标签,默认为删除的评论,添加标签
//  第三 重新设计表
func SearchService(c *gin.Context) *serialize.Response{
	// 最新,最热门,all
	// 标签,课程标题模糊搜索
	start,size 	:= service.PagingQuery(c)

	key 		:= "%"+ c.Query("key")+"%"		// 查询字段
	order 		:= c.Query("order")				// 排序字段

	searchSql 	:= ""

	// select count(*) from curriculums as c where c.c_id in (select ct.c_id from crriculu_tag as ct where ct.t_id in (select t_id from tags where t_name like '%r%')) or c.c_name like '%r%'
	totalSql := "select " +
					"count(1) as total " +
				"from " +
					"(select " +
					"c.c_id " +
				"from " +
					"curriculums as c join crriculu_tag as ct on c.c_id = ct.c_id join tags as t on t.t_id = ct.t_id " +
				"where " +
					"c.delete_at is null and " +
					"c.c_name like ? or t.t_name like ? group by c.c_id ) as count"




	if order == "new" {					// 最新 字段创建时间查询
		// select c.c_id,c.u_id,c.c_name,c.price,c.c_image,c.create_at,avg(cc.number) as number from curriculum_comments as cc join curriculums as c on c.c_id = cc.c_id join crriculu_tag as ct on c.c_id = ct.c_id join tags as t on ct.t_id = t.t_id where c.c_name like '%re%' or t.t_name like '%re%' and c.create_at is null group by c.c_id order by c.create_at asc limit 0,5"
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
					"c.c_name like ? or t.t_name like ? " +	 // 不能使用gorm的字段拼接,只能手动 去除 繁殖sql注入
					"and c.delete_at is null " +
					"group by c.c_id order by c.create_at asc limit ?,?"

	}else if order == "hot" {			// 最热门 学习人数最多的排序 不统计没有在订单表中的数据
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
					"c.c_name like ? or t.t_name like ? " +
				"and c.delete_at is null " +
				"group by c.c_id " +
				"order by count desc " +
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
					"c.c_name like ? or t.t_name like ? and c.delete_at is null " +
				"group by c.c_id limit ?,?"

	}


	keys := []string{key,order,string(start),string(size)}
	cachedata := service.GetCacheTypeStr(keys)
	if cachedata != "" {
		return serialize.Res(cachedata,"")
	}

	var data SearchServiceData
	DB.DB.Raw(searchSql,key,key,start,size).Scan(&data.Result)

	// 查询数据的总和
	DB.DB.Raw(totalSql,key,key).Count(&data.Total)


	service.SetCacheTypeStr(keys,data,0)


	return serialize.Res(data,"")
}

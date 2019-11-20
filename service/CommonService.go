package service

import (
	"demos/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

func pagingQuery(c *gin.Context)(start,end int){
	var err error
	undetermined1 := c.DefaultQuery("page","0")
	undetermined2 := c.DefaultQuery("size","20")

	page,err := strconv.Atoi(undetermined1)
	if (err != nil || page < 0) {
		page = 0
	}

	size ,err := strconv.Atoi(undetermined2)
	if err != nil {
		size = 20
	}else if (size <=0 ||size >= 50) {
		size = 30
	}

	start = page*size
	// end   = page*size+size
	end = size				// 貌似limit时返回 从几个开始的隔开多少个,而不是中间的部分
	return start,end
}



func VerifyUser(c *gin.Context)(model.User,bool){
	s := sessions.Default(c)
	user := s.Get("user")
	if user != nil{
		if chechuser, ok := user.(model.User); ok {

			// 再次验证 user_id 是否存在和相等
			//uid := s.Get("user_id")
			//if string(chechuser.ID) == uid {
			//	return chechuser,true
			//}

			return chechuser,true
		}
	}
	return model.User{},false
}




func GetUserId(c *gin.Context) interface{}{
	s := sessions.Default(c)
	return s.Get("user_id")
}
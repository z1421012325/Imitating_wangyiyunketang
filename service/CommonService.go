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
	end   = page*size+size

	return start,end
}



func VerifyUser(c *gin.Context)(model.User,bool){
	s := sessions.Default(c)
	user := s.Get("user")
	if user != nil{
		if chechuser, ok := user.(*model.User); ok {
			return *chechuser,true
		}
	}
	return model.User{},false
}







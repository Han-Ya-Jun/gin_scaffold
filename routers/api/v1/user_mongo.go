package v1

import (
	models "github.com/Han-Ya-Jun/gin_scaffold/models/mongo"
	"github.com/Han-Ya-Jun/gin_scaffold/pkg/app"
	"github.com/Han-Ya-Jun/gin_scaffold/pkg/e"
	"github.com/Han-Ya-Jun/gin_scaffold/service/user_service"
	_ "github.com/Han-Ya-Jun/gin_scaffold/service/user_service"
	mlog "github.com/e421083458/golang_common/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
* @Author:hanyajun
* @Date:2019/7/11 17:17
* @Name:v1
* @Function: user api
 */

// @title gin_scaffold
// @version 1.0
// @description add user api
// @Summary users add
// @Produce  json
// @Param user body   models.User true "user"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /gin_scaffold/api/v1/mongo/user [post]
func AddUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var user *models.User
	if err := c.ShouldBindJSON(&user); err == nil {
		u := user_service.UserService{User: user}
		err := u.Add()
		if err != nil {
			mlog.Error("add user err:%v", err)
			appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		} else {
			appG.Response(http.StatusOK, e.SUCCESS, nil)
		}

	} else {
		mlog.Error("add user err:%v", err)
		appG.Response(http.StatusOK, e.ERROR, err)
	}

}

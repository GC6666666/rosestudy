package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rose/internal/dto"
)

func userDetail(c *gin.Context) {
	r := new(dto.UserDetailReq)

	if err := c.Bind(r); err != nil {
		return
	}

	ret, err := svc.UserDetail(c.Request.Context(), r.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, ret)
}

func userCreate(c *gin.Context) {
	r := new(dto.UserCreateReq)
	if err := c.BindJSON(r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ret, err := svc.UserCreate(c.Request.Context(), r.UserName)
	if err != nil {
		// panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database insert error!"})
	}
	c.JSON(http.StatusOK, ret)
}

func userList(c *gin.Context) {
	r := new(dto.UserListReq)
	if err := c.Bind(r); err != nil {
		return
	}
	ret, err := svc.UserList(c.Request.Context(), r.UserIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, ret)

}

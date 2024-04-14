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
